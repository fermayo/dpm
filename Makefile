PKG_NAME := github.com/JPZ13/dpm
GO := docker run -it --rm -v ${PWD}:/go/src/$(PKG_NAME) -w /go/src/$(PKG_NAME) -e GOOS -e GOARCH golang:1.7 go
GLIDE := docker run -it --rm -v ${PWD}:/run/context -w /run/context dockerepo/glide

.PHONY: all clean binaries linux-binary mac-binary fmt glide-init glide-install glide-update

all: binaries

clean:
	rm -fr build/

binaries: linux-binary mac-binary

linux-binary:
	GOOS=linux GOARCH=amd64 $(GO) build -v -o build/dpm-Linux-x86_64

mac-binary:
	GOOS=darwin GOARCH=amd64 $(GO) build -v -o build/dpm-Darwin-x86_64

fmt:
	$(GO) fmt ./...

glide-init:
	$(GLIDE) init

glide-install:
	$(GLIDE) install

glide-update:
	$(GLIDE) update
