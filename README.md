# Windows

Download protoc-win32.zip from "https://developers.google.com/protocol-buffers/docs/downloads".

# Install

    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    go get -u github.com/golang/protobuf/proto
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest


# Commands

    go mod init github.com/aldoclg/grpc
    go mod tidy

    $GOPATH/bin/protoc.exe -I=proto/ --go_opt=paths=source_relative --go_out=package/proto/v1/message proto/message.proto
    $GOPATH/bin/protoc.exe -I=proto/ --go_opt=paths=source_relative --go_out=package/proto/v1/servive proto/service.proto
    $GOPATH/bin/protoc.exe -I=proto/ --go-grpc_opt=paths=source_relative --go-grpc_out=package/proto/v1/service proto/service.proto

    go run cmd/main.go
```shell
grpcurl -import-path proto -proto proto/service.proto -plaintext -d @ localhost:50051 greeting.v1.GreeterService/Greet <<EOM
{
    "msg": {
        "greeting": 0,
        "name": "yourname"
    }
}
EOM
```

