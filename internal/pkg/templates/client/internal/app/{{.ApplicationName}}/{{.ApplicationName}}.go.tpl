package {{.ApplicationName}}

import (
	"context"
	client "{{.ModuleName}}/internal/client/helloworld"
	"{{.ModuleName}}/pkg/api/v1/helloworld"
)

func Run() error {
	client.Init()
	client.Client.SayHello(context.Background(), &helloworld.HelloRequest{})
	return nil
}
