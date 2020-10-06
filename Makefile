.PHONY: all golint vet fmt test coverage scan build linux macos windows clean
BUILT_HASH=$(shell git rev-parse --short HEAD)
BUILT_VERSION=v1.0.2
LDFLAGS=-ldflags "-X main.ApplicationVersion=${BUILT_VERSION} -w -s"
TRAVISBUILD?=off
all: clean linting build

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

linting: golintci card

golintci:
	golangci-lint run ./...

card:
	goreportcard-cli -v -t 100

build: linux macos windows

linux: linux32 linux64 linuxarm linuxarm64 linuxpi

windows: windows32 windows64

linux64:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ./builds/linux-amd64/go-transip-dyndns
	cp README.md ./builds/linux-amd64/
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./builds/linux-amd64/go-transip-dyndns
endif
	@cd ./builds/linux-amd64 && tar -jcf ../../go-transip-dyndns-linux-amd64-${BUILT_VERSION}.tbz2 go-transip-dyndns README.md
	@cd ./builds/linux-amd64 && tar -zcf ../../go-transip-dyndns-linux-amd64-${BUILT_VERSION}.tgz go-transip-dyndns README.md
	cp docker/Dockerfile ./builds/linux-amd64
	cp docker/docker-compose.yml ./builds/linux-amd64


linux32:
	env GOOS=linux GOARCH=386 go build ${LDFLAGS} -o ./builds/linux-386/go-transip-dyndns
	cp README.md ./builds/linux-386/
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./builds/linux-386/go-transip-dyndns
endif
	@cd ./builds/linux-386 && tar -jcf ../../go-transip-dyndns-linux-386-${BUILT_VERSION}.tbz2 go-transip-dyndns README.md
	@cd ./builds/linux-386 && tar -zcf ../../go-transip-dyndns-linux-386-${BUILT_VERSION}.tgz go-transip-dyndns README.md
	cp docker/Dockerfile ./builds/linux-386
	cp docker/docker-compose.yml ./builds/linux-386


linuxarm:
	env GOOS=linux GOARCH=arm go build ${LDFLAGS} -o ./builds/linux-arm/go-transip-dyndns
	cp README.md ./builds/linux-arm/
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./builds/linux-arm/go-transip-dyndns
endif
	@cd ./builds/linux-arm && tar -jcf ../../go-transip-dyndns-linux-arm-${BUILT_VERSION}.tbz2 go-transip-dyndns README.md
	@cd ./builds/linux-arm && tar -zcf ../../go-transip-dyndns-linux-arm-${BUILT_VERSION}.tgz go-transip-dyndns README.md
	cp docker/Dockerfile ./builds/linux-arm
	cp docker/docker-compose.yml ./builds/linux-arm


linuxpi:
	env GOOS=linux GOARCH=arm GOARM=5 go build ${LDFLAGS} -o ./builds/linux-pi/go-transip-dyndns
	cp README.md ./builds/linux-pi/
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./builds/linux-pi/go-transip-dyndns
endif
	@cd ./builds/linux-pi && tar -jcf ../../go-transip-dyndns-linux-arm-pi-${BUILT_VERSION}.tbz2 go-transip-dyndns README.md
	@cd ./builds/linux-pi && tar -zcf ../../go-transip-dyndns-linux-arm-pi-${BUILT_VERSION}.tgz go-transip-dyndns README.md
	cp docker/Dockerfile ./builds/linux-pi
	cp docker/docker-compose.yml ./builds/linux-pi


linuxarm64:
	env GOOS=linux GOARCH=arm64 go build ${LDFLAGS} -o ./builds/linux-arm64/go-transip-dyndns
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./builds/linux-arm64/go-transip-dyndns
endif
	cp README.md ./builds/linux-arm64/
	@cd ./builds/linux-arm64 && tar -jcf ../../go-transip-dyndns-linux-arm64-${BUILT_VERSION}.tbz2 go-transip-dyndns README.md
	@cd ./builds/linux-arm64 && tar -zcf ../../go-transip-dyndns-linux-arm64-${BUILT_VERSION}.tgz go-transip-dyndns README.md
	cp docker/Dockerfile ./builds/linux-arm64
	cp docker/docker-compose.yml ./builds/linux-arm64

macos:
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ./builds/macos/go-transip-dyndns
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./builds/macos/go-transip-dyndns
endif
	cp README.md ./builds/macos/
	@cd builds/macos/ && zip ../../go-transip-dyndns-macos-amd64-${BUILT_VERSION}.zip go-transip-dyndns README.md
	cp docker/Dockerfile ./builds/macos
	cp docker/docker-compose.yml ./builds/macos

windows64:
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ./builds/windows-64/go-transip-dyndns.exe
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./builds/windows-64/go-transip-dyndns.exe
endif
	cp README.md ./builds/windows-64/
	@cd builds/windows-64/ && zip ../../go-transip-dyndns-windows-amd64-${BUILT_VERSION}.zip go-transip-dyndns.exe README.md


windows32:
	env GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ./builds/windows-32/go-transip-dyndns.exe
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./builds/windows-32/go-transip-dyndns.exe
endif
	cp README.md ./builds/windows-32/
	@cd builds/windows-32/ && zip ../../go-transip-dyndns-windows-386-${BUILT_VERSION}.zip go-transip-dyndns.exe README.md
