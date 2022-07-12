package main

import (
	"fmt"
	"golang.org/x/tools/go/packages"
)

func main() {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedExportFile |
			packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedTypesInfo |
			packages.NeedTypesSizes |
			packages.NeedModule |
			packages.NeedEmbedFiles |
			packages.NeedEmbedPatterns,
		BuildFlags: []string{"-buildvcs=false"},
		Overlay:    map[string][]byte{},
	}, "./...")
	if err != nil {
		panic(err)
	}
	for _, p := range pkgs {
		fmt.Println(p.Name)
	}
}
