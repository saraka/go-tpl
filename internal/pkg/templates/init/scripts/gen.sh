cd `dirname $0`/..
protoc --proto_path=. --go-grpc_out=. --go_out=. $(find api -name *.proto)