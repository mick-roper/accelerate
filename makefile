.PHONY: all

all: clean install test build-all

clean: 
	rm -rf bin

install:
	cd app && go get -v ./...

test:
	cd app && go test ./...

build-all: build-macos

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/app ./app/main.go