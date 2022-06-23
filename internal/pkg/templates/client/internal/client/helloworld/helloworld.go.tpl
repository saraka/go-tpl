package helloworld

import (
	"crypto/tls"
	"{{.ModuleName}}/pkg/api/v1/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var Client helloworld.GreeterClient

func Init() error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		return err
	}
	Client = helloworld.NewGreeterClient(conn)
	return nil
}
