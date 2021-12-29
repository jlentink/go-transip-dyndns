.PHONY: all golint vet fmt test coverage scan build linux macos windows clean
BUILT_HASH=$(shell git rev-parse --short HEAD)
BUILT_VERSION=v1.0.2
LDFLAGS=-ldflags "-X main.ApplicationVersion=${BUILT_VERSION} -w -s"
TRAVISBUILD?=off
all: clean linting test-release
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

clean:
	@-rm test-report.out
	@-rm coverage.out
	@-rm *.zip
	@-rm *.tbz2
	@-rm *.tgz
	@-rm go-transip-dyndns
	@-rm -rf builds/linux*
	@-rm -rf builds/macos*
	@-rm -rf builds/windows*
	@-rm -rf dist/*

linting: golintci card

golintci:
	golangci-lint run ./...

card:
	goreportcard-cli -v -t 100

build:
	goreleaser build --skip-validate --rm-dist

release:
	goreleaser release --rm-dist

test-release:
	goreleaser release --rm-dist --skip-validate --skip-publish

test-docker:
	docker run --rm -v ${ROOT_DIR}/go-transip-dyndns.toml:/etc/go-transip-dyndns.toml jlentink/go-transip-dyndns:latest

test-docker-shell:
	 docker run -it --rm -v ${ROOT_DIR}/go-transip-dyndns.toml:/etc/go-transip-dyndns.toml jlentink/go-transip-dyndns:latest /bin/sh