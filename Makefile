TIME:=$(shell date -u "+%Y-%m-%dT%H:%M:%S")
HASH:=$(shell git log --format=format:%h -n 1; [ -n "$(shell git status --porcelain)")
TAG:=$(shell git describe --tags --exact-match 2> /dev/null)

all: build

.PHONY: build
build:
	go build -ldflags "-X main.buildTag=$(TAG) -X main.buildTime=$(TIME) -X main.buildHash=$(HASH)" \
		-o imgbb-cli github.com/JohnNON/ImgBB-cli/cmd

.PHONY: linux-amd64
linux-amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.buildTag=$(TAG) -X main.buildTime=$(TIME) -X main.buildHash=$(HASH)" \
		-o imgbb-cli github.com/JohnNON/ImgBB-cli/cmd

.PHONY: linux-arm64
linux-arm64:
	GOOS=linux GOARCH=arm64 go build -ldflags "-X main.buildTag=$(TAG) -X main.buildTime=$(TIME) -X main.buildHash=$(HASH)" \
		-o imgbb-cli github.com/JohnNON/ImgBB-cli/cmd

.PHONY: macos-arm64
macos-arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.buildTag=$(TAG) -X main.buildTime=$(TIME) -X main.buildHash=$(HASH)" \
		-o imgbb-cli github.com/JohnNON/ImgBB-cli/cmd

.PHONY: macos-amd64
macos-amd64:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.buildTag=$(TAG) -X main.buildTime=$(TIME) -X main.buildHash=$(HASH)" \
		-o imgbb-cli github.com/JohnNON/ImgBB-cli/cmd

.PHONY: windows
windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.buildTag=$(TAG) -X main.buildTime=$(TIME) -X main.buildHash=$(HASH)" \
		-o imgbb-cli.exe github.com/JohnNON/ImgBB-cli/cmd
