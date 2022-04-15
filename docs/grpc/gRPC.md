## Install
1. protoc [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)
2. protoc-gen-go and protoc-gen-go-grpc
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Generate gRPC code
```shell
# Move helloworld.proto to ROOT/grpc
# cd ROOT then execute in root path
protoc --go_out=./ --go-grpc_out=./ ./grpc/helloworld.proto
```


## Docs 
https://grpc.io/docs/languages/go/quickstart  
https://github.com/grpc/grpc-go/tree/master/examples  
