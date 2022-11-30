package main

import (
	"context"
	"flag"
	"log"

	"petstore/internal/gateway"
)

var (
	endpoint = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	network  = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
)

func main() {
	flag.Parse()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":8080",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
	}
	if err := gateway.Run(ctx, opts); err != nil {
		log.Fatal(err)
	}
}
