TIME:=$(shell date -u "+%Y-%m-%dT%H:%M:%S")
HASH:=$(shell git log --format=format:%h -n 1; [ -n "$(shell git status --porcelain)")
TAG:=$(shell git describe --tags --exact-match 2> /dev/null)

BIN_DIR = ./bin
BINARY = $(BIN_DIR)/imgbb-cli
SRC = github.com/JohnNON/ImgBB-cli/cmd
LDFLAGS = -ldflags="-s -w -X main.buildTag=$(TAG) -X main.buildTime=$(TIME) -X main.buildHash=$(HASH)"

all: macos windows linux zip
	@echo "complete"

create_dist_dir:
	@mkdir -p $(BIN_DIR)

macos: create_dist_dir
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o $(BINARY)_macos_amd64 $(LDFLAGS) $(SRC)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o $(BINARY)_macos_arm64 $(LDFLAGS) $(SRC)

windows: create_dist_dir
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o $(BINARY)_windows_amd64.exe $(LDFLAGS) $(SRC)
	GOOS=windows GOARCH=arm64 CGO_ENABLED=1 go build -o $(BINARY)_windows_arm64.exe $(LDFLAGS) $(SRC)

linux: create_dist_dir
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-linux-gnu-gcc  go build -o $(BINARY)_linux_amd64 $(LDFLAGS) $(SRC)
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc go build -o $(BINARY)_linux_arm64 $(LDFLAGS) $(SRC)

zip:
	@for f in $(shell ls ${BIN_DIR}); do zip ${BIN_DIR}/$${f}.zip ${BIN_DIR}/$${f} && rm ${BIN_DIR}/$${f}; done

clean:
	rm -rfd $(BIN_DIR)
