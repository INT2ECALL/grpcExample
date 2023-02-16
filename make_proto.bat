protoc.exe --go_out=. user.proto
protoc -I. --go_grpc_out=./ user.proto