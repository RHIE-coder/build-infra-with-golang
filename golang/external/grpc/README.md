# gRPC

## [Install]

### - protobuf (`protoc`)

https://github.com/protocolbuffers/protobuf/releases/tag/v23.1

```sh
wget https://github.com/protocolbuffers/protobuf/releases/download/v23.1/protoc-23.1-linux-x86_64.zip
unzip protoc-23.1-linux-x86_64.zip -d /usr/local
unzip protoc-23.1-linux-x86_64.zip 
```

## [SPEC]

### - gRPC Documentation

https://grpc.io/docs/what-is-grpc/introduction/

### - Protocol Buffers Version 3 Language Specification

https://protobuf.dev/reference/protobuf/proto3-spec/


## [HTTP/2]

### - 지원 브라우저 목록 확인

https://caniuse.com/http2

### - 주요 구현 라이브러리 목록 확인

https://github.com/httpwg/http2-spec/wiki/Implementations

## [start-app]

 - go.mod `module hello/protobuf`
 - hello/protobuf/helloworld.proto

```proto
syntax = "proto3";

package helloworld;

option go_package = "golang/external/grpc";

service Greeter {
    rpc SayHello (HelloRequest) returns (HelloReply) {}
    rpc SayHelloAgain (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
    string name = 1;
}

message HelloReply {
    string message = 1;
}
```

 - generate gRPC code

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto
```