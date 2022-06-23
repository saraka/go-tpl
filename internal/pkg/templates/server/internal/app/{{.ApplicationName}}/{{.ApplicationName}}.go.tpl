package {{.ApplicationName}}

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	server "{{.ModuleName}}/internal/server/helloworld"
	"{{.ModuleName}}/pkg/api/v1/helloworld"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func Run() error {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, server.Server)
	return s.Serve(lis)
}
