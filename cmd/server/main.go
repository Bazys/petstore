package main

import (
	"context"
	"flag"
	"log"

	"petstore/internal/server"
)

var (
	addr    = flag.String("addr", ":9090", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	if err := server.Run(ctx, *network, *addr); err != nil {
		log.Fatal(err)
	}
}
