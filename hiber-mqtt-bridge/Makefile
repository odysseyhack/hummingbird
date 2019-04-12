EXECUTABLE = hummingbird-hiber-bridge

VARIANTS = linux-arm
ALL_FILES = $(wildcard *.go)

all: format build
.PHONY: all

build/$(EXECUTABLE)-linux-arm: ${ALL_FILES}
	go get golang.org/x/sys/unix
	GOARCH=arm GOOS=linux GOARM=5 go build -ldflags=-s -o "$@" || true

format:
	go fmt

build: $(foreach VARIANT,$(VARIANTS),build/$(EXECUTABLE)-$(VARIANT))

.PHONY: build

test: $(wildcard *.go)
	go test ./...