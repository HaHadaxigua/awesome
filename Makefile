.PHONY: all build compile clean

BUILDTIME ?= $(shell date +%Y-%m-%d_%I:%M:%S)
GITCOMMIT ?= $(shell git rev-parse -q HEAD)
ifeq ($(CI_PIPELINE_ID),)
	BUILDNUMER := private
else
	BUILDNUMER := $(CI_PIPELINE_ID)
endif
VERSION ?= $(shell git describe --tags --always --dirty)

LDFLAGS = -extldflags \
		  -static \
		  -X "main.Version=$(VERSION)" \
		  -X "main.BuildTime=$(BUILDTIME)" \
		  -X "main.GitCommit=$(GITCOMMIT)" \
		  -X "main.BuildNumber=$(BUILDNUMER)"


test_compile:
	cd ./tencent
	GOOS=linux GOARCH=amd64 go build -o tencent/tencent ./tencent/main.go

test_yeagi:
	cd yaegi/cmd;
	go generate;
	GOOS=linux GOARCH=amd64 go build -o yaegi/bin/yeagi main.go

test_yeagi_try:
	cd yaegi/linux_try;\
    GOOS=linux GOARCH=amd64 go build -o yaegi/bin/yeagi_linux main.go

test_wasm:
	cd wasm;
	cargo build demo.rs