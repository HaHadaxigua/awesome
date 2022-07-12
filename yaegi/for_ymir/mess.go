package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"testing/fstest"
)

func buildFS(path string) (fstest.MapFS, error) {
	pkgs, err := loadPackages(path)
	if err != nil {
		return nil, err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		logrus.Errorf("failed to get absoluate path of %s", path)
		return nil, err
	}
	fs := make(fstest.MapFS)
	for _, pkg := range pkgs {
		for _, fp := range pkg.GoFiles {
			// fill fs
			buf, err := os.ReadFile(fp)
			if err != nil {
				return nil, err
			}
			fspath := filepath.Join(pkg.Module.Path, strings.TrimLeft(fp, absPath))
			fspath = filepath.Join(fakeFSPrefix, fspath)
			fs[fspath] = &fstest.MapFile{
				Data: buf,
				Mode: 0755,
				Sys:  fp,
			}
		}
	}
	return fs, nil
}

func loadPackages(path string) ([]*packages.Package, error) {
	if util.IsDirNotExist(path) {
		return nil, fmt.Errorf("project %s not found", path)
	}

	originWD, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to load current work dir: %v", err)
	}
	defer os.Chdir(originWD)

	if err := os.Chdir(path); err != nil {
		logrus.Errorf("failed to switch workdir to %s", path)
		return nil, err
	}

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
		logrus.Errorf("failed to load all packages from %s", path)
		return nil, err
	}

	for _, pkg := range pkgs {
		if len(pkg.Errors) != 0 {
			packages.PrintErrors([]*packages.Package{pkg})
			return nil, fmt.Errorf("invliad go package: %s, location: [%s]", pkg.PkgPath, strings.Join(pkg.GoFiles, ","))
		}
	}
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found or it is an invalid go project")
	}
	return pkgs, nil
}
