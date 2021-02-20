.PHONY: all test build clean

all: test build

build: 
	mkdir -p build
	#go build -o build -tags real ./...
	go build -o build ./...

build-arm64:
	mkdir -p buildarm
	GOOS=linux GOARCH=arm64 go build -o buildarm -ldflags="-s -w" ./...


build-dev:
	mkdir -p build
	GOOS=linux go build -ldflags="-s -w" -o build ./...
	chmod 755 build/microservice
	chmod 755 build/uid_entrypoint.sh

test:
	go test -v -coverprofile=tests/results/cover.out -tags fake ./...

cover:
	go tool cover -html=tests/results/cover.out -o tests/results/cover.html

clean:
	rm -rf build/*
	rm -rf buildarm/*
	go clean ./...

container:
	podman build -t quay.io/14west/agoracxp-couchbase-interface:1.15.6 .

push:
	podman push quay.io/14west/agoracxp-couchbase-interface:1.15.6 
