APP=jyksServer
VERSION=0.0.1

.PHONY: help all build windows linux

help:
	@echo "使用 make 编译程序"
all:build windows-amd64 linux-arm linux-amd64
build:
	@go build -o ${APP}
windows-amd64:
	@go env -w GOOS=windows
	@go env -w GOARCH=amd64
	@go build -o ${APP}-windows-amd64-${VERSION}
linux-arm:
	@go env -w GOOS=linux
	@go env -w GOARCH=arm
	@go build -o ${APP}-linux-arm-${VERSION}
linux-amd64:
	@go env -w GOOS=linux
	@go env -w GOARCH=amd64
	@go build -o ${APP}-linux-amd64-${VERSION}