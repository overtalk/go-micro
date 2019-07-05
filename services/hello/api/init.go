package api

import (
	"context"

	proto "github.com/qinhan-shu/go-micro/services/hello/proto"
)

// Greeter :
type Greeter struct {
}

// Hello : 实现服务的 Hello 方法
func (g *Greeter) Hello(ctx context.Context, in *proto.HelloRequest, out *proto.HelloResponse) error {
	out.Greeting = "Hello " + in.Name
	return nil
}
