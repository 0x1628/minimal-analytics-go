PROJECT="minimal-analytics-go"

default:
	echo "hello world"

install:
	mkdir -p gen && protoc -I=./proto --go_out=paths=source_relative:./gen ./proto/events.proto

.PHONY: default install