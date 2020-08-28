PROJECT="minimal-analytics-go"

default:
	echo "hello world"

install:
	mkdir -p gen && protoc -I=./proto --go_out=paths=source_relative:./gen ./proto/maevents.proto

build:
	go build -o ma

dev: build
	./ma

.PHONY: default install build