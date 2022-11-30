package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"petstore/pkg/pb"

	"google.golang.org/grpc"
)

func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			fmt.Printf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	pb.RegisterPetstoreServer(s, newPetstoreServer())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
		log.Println("Shutting down the http gateway server")
	}()

	fmt.Printf("Starting listening server at %s\n", address)
	return s.Serve(l)
}
