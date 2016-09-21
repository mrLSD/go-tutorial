#
# Makefile
# @author Evgeny Ukhanov <mrlsd@ya.ru>
#

.PHONY: run, test, build, fmt, docker-build, docker-rmi, docker-bld, install-go

default: run

run:
	@go run main.go

test:
	@go test ./...

build:
	@go build

fmt:
	@gofmt -w -l .
