package saml

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/crewjam/saml/samlsp"
	"github.com/dchest/uniuri"
	"github.com/kr/pretty"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

var links = map[string]Link{}

// Link represents a short link
type Link struct {
	ShortLink string
	Target    string
	Owner     string
}

// CreateLink handles requests to create links
func CreateLink(c web.C, w http.ResponseWriter, r *http.Request) {
	account := r.Header.Get("X-Remote-User")
	l := Link{
		ShortLink: uniuri.New(),
		Target:    r.FormValue("t"),
		Owner:     account,
	}
	links[l.ShortLink] = l

	fmt.Fprintf(w, "%s\n", l.ShortLink)
	return
}

// ServeLink handles requests to redirect to a link
func ServeLink(c web.C, w http.ResponseWriter, r *http.Request) {
	l, ok := links[strings.TrimPrefix(r.URL.Path, "/")]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, l.Target, http.StatusFound)
	return
}

// ListLinks returns a list of the current user's links
func ListLinks(c web.C, w http.ResponseWriter, r *http.Request) {
	account := r.Header.Get("X-Remote-User")
	for _, l := range links {
		if l.Owner == account {
			fmt.Fprintf(w, "%s\n", l.ShortLink)
		}
	}
}

var (
	serviceKey, _  = os.ReadFile("service.key")
	serviceCert, _ = os.ReadFile("service.cert")
)

func TestService(t *testing.T) {
	rootURLstr := flag.String("url", "http://localhost:8000", "The base URL of this service")
	idpMetadataURLstr := flag.String("idp", "http://localhost:8000/idp/metadata", "The metadata URL for the IDP")
	flag.Parse()

	keyPair, err := tls.X509KeyPair(serviceCert, serviceKey)
	if err != nil {
		panic(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err) // TODO handle error
	}

	idpMetadataURL, err := url.Parse(*idpMetadataURLstr)
	if err != nil {
		panic(err) // TODO handle error
	}

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		panic(err) // TODO handle error
	}

	rootURL, err := url.Parse(*rootURLstr)
	if err != nil {
		panic(err) // TODO handle error
	}

	samlSP, err := samlsp.New(samlsp.Options{
		URL:               *rootURL,
		Key:               keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:       keyPair.Leaf,
		AllowIDPInitiated: true,
		IDPMetadata:       idpMetadata,
	})
	if err != nil {
		panic(err) // TODO handle error
	}

	// register with the service provider
	spMetadataBuf, _ := xml.MarshalIndent(samlSP.ServiceProvider.Metadata(), "", "  ")

	spURL := *idpMetadataURL
	spURL.Path = "/services/sp"
	http.Post(spURL.String(), "text/xml", bytes.NewReader(spMetadataBuf))

	goji.Handle("/saml/*", samlSP)

	authMux := web.New()
	authMux.Use(samlSP.RequireAccount)
	authMux.Get("/whoami", func(w http.ResponseWriter, r *http.Request) {
		pretty.Fprintf(w, "%# v", r)
	})
	authMux.Post("/", CreateLink)
	authMux.Get("/", ListLinks)

	goji.Handle("/*", authMux)
	goji.Get("/:link", ServeLink)

	goji.Serve()
}
