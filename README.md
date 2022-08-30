### grpc-cep

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2


protoc --go-grpc_out=require_unimplemented_servers=true:. -I=%SRC_DIR% --go_out=%SRC_DIR%\protobuf %SRC_DIR%\protobuf\cep.proto