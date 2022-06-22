package {{.ProjectName}}

import (
	"context"
	"flag"
	"fmt"
	"{{.ModuleName}}/pkg/api/v1/helloworld"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s server) SayHello(context.Context, *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, nil
}

func Run() error {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	return s.Serve(lis)
}
