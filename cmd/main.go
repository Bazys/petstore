package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"petstore/internal/gateway"
	"petstore/internal/server"

	"golang.org/x/sync/errgroup"
)

var (
	addr     = flag.String("addr", ":9090", "endpoint of the gRPC service")
	endpoint = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	network  = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
)

func main() {
	flag.Parse()
	ctx := context.Background()
	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	termSignal := make(chan os.Signal, 1)
	defer close(termSignal)
	signal.Notify(termSignal, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	group, ctx := errgroup.WithContext(ctxCancel)

	stopCh := make(chan struct{})

	group.Go(func() error {
		sig := <-termSignal
		log.Println("server: shutting down... reason:", sig.String())
		close(stopCh)
		cancel()
		return nil
	})

	opts := gateway.Options{
		Addr: ":8080",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
	}
	group.Go(func() error {
		return server.Run(ctxCancel, *network, *addr)
	})

	group.Go(func() error {
		return gateway.Run(ctxCancel, opts)
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("application is failed: %s", err)
	}

}
