BINARY = bin/daemon
BUILD_FLAGS = -ldflags="-s -w" 
CMD=cmd/main.go
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
IMAGE= codebaker/$(WORKINGDIR):$(TAG)
TAG=$(shell git describe)
UNAME= $(shell uname | tr '[:upper:]' '[:lower:]')
WORKINGDIR= $(shell pwd | rev | cut -d'/' -f 1 | rev)

build:
	#building app static binary ...
	env GOOS=linux CGO_ENABLED=0 go build -a -installsuffix nocgo $(BUILD_FLAGS) -o $(BINARY) ./server/.

certificates:
	#1
	openssl genrsa -out cert/server.key 2048

	#2
	openssl req -new -x509 -sha256 -key cert/server.key \
	-out cert/server.crt -days 3650

	#3
	openssl req -new -sha256 -key cert/server.key -out cert/server.csr

	#4
	openssl x509 -req -sha256 -in cert/server.csr \
	-signkey cert/server.key -out cert/server.crt -days 3650

dependecies:
	#installing dependecies ...
	go mod tidy

docs:
	#generating docs ...
	docker run --rm \
  	-v $(shell pwd)/proto/docs:/out \
	-v $(shell pwd)/proto/cache:/protos \
  	pseudomuto/protoc-gen-doc --doc_opt=markdown,docs.md

image:
	#building docker image ...
	docker build -t $(IMAGE) .

install:
	#install the golang command line interface
	go install ./cli/quavoctl

protoc:
	#generate protocol buffer(.proto) files ...
	protoc \
		-I proto/cache proto/cache/cache.proto \
		--go_out=plugins=grpc:models/cache \


start:
	#run quavo deamon
	./bin/daemon

