package helloworld

import (
	"context"
	"{{.ModuleName}}/pkg/api/v1/helloworld"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s server) SayHello(context.Context, *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, nil
}

var Server helloworld.GreeterServer = &server{}
