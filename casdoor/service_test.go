package casdoor

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	saml2 "github.com/russellhaering/gosaml2"
	"github.com/russellhaering/gosaml2/types"
	dsig "github.com/russellhaering/goxmldsig"
)

func TestCasdoorSAML(t *testing.T) {
	res, err := http.Get("http://localhost:8001/api/saml/metadata?application=admin/app-built-in")
	//res, err := http.Get("http://localhost:7001/api/saml/metadata?application=admin/application_p1d3ur")
	if err != nil {
		panic(err)
	}

	rawMetadata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	metadata := &types.EntityDescriptor{}
	err = xml.Unmarshal(rawMetadata, metadata)
	if err != nil {
		panic(err)
	}

	certStore := dsig.MemoryX509CertificateStore{
		Roots: []*x509.Certificate{},
	}

	for _, kd := range metadata.IDPSSODescriptor.KeyDescriptors {
		for idx, xcert := range kd.KeyInfo.X509Data.X509Certificates {
			if xcert.Data == "" {
				panic(fmt.Errorf("metadata certificate(%d) must not be empty", idx))
			}
			certData, err := base64.StdEncoding.DecodeString(xcert.Data)
			if err != nil {
				panic(err)
			}

			idpCert, err := x509.ParseCertificate(certData)
			if err != nil {
				panic(err)
			}

			certStore.Roots = append(certStore.Roots, idpCert)
		}
	}

	randomKeyStore := dsig.RandomKeyStoreForTest()

	sp := &saml2.SAMLServiceProvider{
		IdentityProviderSSOURL:      metadata.IDPSSODescriptor.SingleSignOnServices[0].Location,
		IdentityProviderIssuer:      metadata.EntityID,
		ServiceProviderIssuer:       "http://localhost:6900/acs/example",
		AssertionConsumerServiceURL: "http://localhost:6900/v1/_saml_callback",
		SignAuthnRequests:           true,
		AudienceURI:                 "http://localhost:6900/saml/acs/example",
		IDPCertificateStore:         &certStore,
		SPKeyStore:                  randomKeyStore,
	}

	http.HandleFunc("/v1/_saml_callback", func(rw http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		samlReponse := req.URL.Query().Get("SAMLResponse")
		assertionInfo, err := sp.RetrieveAssertionInfo(samlReponse)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusForbidden)
			return
		}
		fmt.Println(assertionInfo)
		if assertionInfo.WarningInfo.InvalidTime {
			fmt.Println("here12:", assertionInfo.WarningInfo.InvalidTime)
			rw.WriteHeader(http.StatusForbidden)
			return
		}

		if assertionInfo.WarningInfo.NotInAudience {
			fmt.Println(assertionInfo)
			fmt.Println("here13:", assertionInfo.WarningInfo.NotInAudience)
			rw.WriteHeader(http.StatusForbidden)
			return
		}

		fmt.Fprintf(rw, "NameID: %s\n", assertionInfo.NameID)

		fmt.Fprintf(rw, "Assertions:\n")

		for key, val := range assertionInfo.Values {
			fmt.Fprintf(rw, "  %s: %+v\n", key, val)
		}
		fmt.Println(assertionInfo.Values.Get("FirstName"))
		fmt.Fprintf(rw, "\n")

		fmt.Fprintf(rw, "Warnings:\n")
		fmt.Fprintf(rw, "%+v\n", assertionInfo.WarningInfo)
	})

	println("Visit this URL To Authenticate:")
	authURL, err := sp.BuildAuthURL("")
	if err != nil {
		panic(err)
	}

	println(authURL)

	println("Supply:")
	fmt.Printf("  SP ACS URL      : %s\n", sp.AssertionConsumerServiceURL)

	err = http.ListenAndServe(":6900", nil)
	if err != nil {
		panic(err)
	}
}
