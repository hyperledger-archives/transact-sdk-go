# Building Protocol Buffers
This repository uses generated code via Google's `protoc-gen-go`
utility, which generates Go code from proto specification.

The source protobufs are found in [protos/](protos/). The transaction
generics `batch.proto` and `transaction.proto` are sourced from the
[hyperledger/transact](https://github.com/hyperledger/transact/tree/master/libtransact/protos) repository, whereas `sabre_payload.proto` is
sourced from [Cargill/splinter](https://github.com/Cargill/splinter/blob/master/examples/gameroom/gameroom-app/sabre_proto/sabre_payload.proto), or alternatively [hyperledger/transact-sdk-go](https://github.com/hyperledger/transact-sdk-javascript/tree/master/protos).

# Build
**Requirements**
1. Protobuf Compiler - the protoc binary must be accesible on the build
`$PATH`. Download the binary from the protocolbuffers/protobuf
repository [releases](https://github.com/protocolbuffers/protobuf/releases).
2. Install Golang's `protoc-gen-go` protocol buffers plugin, which is
maintained as a dependency in go.mod.
    ```shell
    go get google.golang.org/protobuf/cmd/protoc-gen-go
    go install google.golang.org/protobuf/cmd/protoc-gen-go
    ```

**Build Command**
```shell
./scripts/compile_protos.sh
```
