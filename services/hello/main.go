package main

import (
	"fmt"
	"os"

	// "context"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	api "github.com/qinhan-shu/go-micro/services/hello/api"
	proto "github.com/qinhan-shu/go-micro/services/hello/proto"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),

		// Setup some flags. Specify --run_client to run the client

		// Add runtime flags
		// We could do this below too
		micro.Flags(
			cli.BoolFlag{
				Name:  "run_client",
				Usage: "Launch the client",
			},
			cli.StringFlag{
				Name:  "str",
				Usage: "test string",
			},
		),
	)

	service.Init(
		// Add runtime action
		// We could actually do this above
		micro.Action(func(c *cli.Context) {
			flag := false
			if c.Bool("run_client") {
				fmt.Println(" run_client yes ")
				flag = true
			}
			if c.String("str") != "" {
				fmt.Println(c.String("str"))
				flag = true
			}
			if flag {
				os.Exit(0)
			}
		}),
	)

	proto.RegisterHelloHandler(service.Server(), new(api.Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
