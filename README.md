# Quavo

Quavo is a proof of concept key-value store that uses protocol buffers via gRPC for server/client communication.

## Getting started
1.  run `make dependecies` to installl dependencies.
2.  run `make build` to build server binary.
3.  run `make start` to run the dev server.
3.  run `make install` to install the CLI tool (quavoclt).

## Basic help
Run `quavoctl --help`

## To generate docs run
1. `docker pull pseudomuto/protoc-gen-doc` to install protoc-gen-doc.
2. `make docs` to automate docs generation.

Documentation [here](proto/docs/docs.md)