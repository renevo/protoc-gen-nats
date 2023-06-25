# Protobuf Generator for NATS Micro Services

This prototype code generator will take protobuf RPC and render services into [NATS](https://nats.io) [micro services](https://pkg.go.dev/github.com/nats-io/nats.go/micro). The reason for the creation of this was to start building a framework on top of the `micro.Service` implementation that is more usable by end users.

See the [examples](./examples/) for the proto files input and output created [hello.nats.go](./examples/hello.nats.go).

This is not yet complete, and has some work left to do.

## Installation

```bash
go install github.com/renevo/protoc-gen-nats@latest
```

## Example Usage

Be sure to have the [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/) installed.

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/renevo/protoc-gen-nats@latest

protoc --go_out=. --go_opt=paths=source_relative --nats_out=. --nats_opt=source_relative ./examples/hello.proto
```
