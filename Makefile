PROJECT="minimal-analytics-go"

default:
	echo "hello world"

install:
	mkdir -p gen && protoc -I=./proto --go_out=paths=source_relative:./gen ./proto/maevents.proto

build:
	go build -o ma-v2

dev:
	go run main.go

.PHONY: default install build dev
