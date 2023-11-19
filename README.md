# AppPC

A simple gRPC server to manage courses.

### Prerequisites

You will need a gRPC client, I recommend [evans](https://github.com/ktr0731/evans).

### Installing

Install evans using go

```
go install github.com/ktr0731/evans@latest
```

And check the version

```
evans --version
```

Finally, install project dependencies locally

```
go mod download
```

### Running

Initiate your server

```
go run cmd/grpc_server/main.go
```

And access your server by the gRPC client

```
evans -r repl
```

## Built With

* [gRPC](https://grpc.io/) - Remote Procedure Call (RPC) framework.
* [Protocol Buffers](https://protobuf.dev/) - Mechanism for serializing structured data.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

