BINARY_NAME=wcgo

SRC_FILES=$(shell find . -name '*.go')

OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m | sed 's/x86_64/amd64/;s/i686/386/;s/aarch64/arm64/')

.PHONY: all
all: build

.PHONY: build_dir
build_dir:
	mkdir -p build

.PHONY: build
build: build_dir $(SRC_FILES)
	GOOS=$(OS) GOARCH=$(ARCH) go build -o build/$(BINARY_NAME) main.go

.PHONY: clean
clean:
	rm -rf build

.PHONY: run
run: build
	./build/$(BINARY_NAME) $(ARGS)

.PHONY: build-linux
build-linux: build_dir $(SRC_FILES)
	GOOS=linux GOARCH=amd64 go build -o build/$(BINARY_NAME)-linux main.go

.PHONY: build-windows
build-windows: build_dir $(SRC_FILES)
	GOOS=windows GOARCH=amd64 go build -o build/$(BINARY_NAME).exe main.go

.PHONY: build-macos
build-macos: build_dir $(SRC_FILES)
	GOOS=darwin GOARCH=amd64 go build -o build/$(BINARY_NAME)-macos main.go

.PHONY: build-arm
build-arm: build_dir $(SRC_FILES)
	GOOS=linux GOARCH=arm go build -o build/$(BINARY_NAME)-arm main.go

.PHONY: build-all
build-all: build-linux build-windows build-macos build-arm
