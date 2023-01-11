package gateway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Endpoint struct {
	Network, Addr string
}

type Options struct {
	// Addr is the address to listen
	Addr string

	// GRPCServer defines an endpoint of a gRPC service
	GRPCServer Endpoint

	// Mux is a list of options to be passed to the gRPC-Gateway multiplexer
	Mux []gwruntime.ServeMuxOption
}

func Run(ctx context.Context, opts Options) error {
	conn, err := dial(ctx, opts.GRPCServer.Network, opts.GRPCServer.Addr)
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.HandleFunc("/healthz", healthzServer(conn))

	gw, err := newGateway(ctx, conn, opts.Mux)
	if err != nil {
		return err
	}

	// r.Handle("/*", gw)
	r.MethodNotAllowed(gw.ServeHTTP)
	r.NotFound(gw.ServeHTTP)

	// chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	// 	fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
	// 	return nil
	// })

	s := &http.Server{
		Addr:    opts.Addr,
		Handler: r,
	}
	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down the http server\n")
		if err := s.Shutdown(context.Background()); err != nil {
			fmt.Errorf("Failed to shutdown http server: %v\n", err)
		}
		if err := conn.Close(); err != nil {
			fmt.Printf("Failed to close a client connection to the gRPC server: %v", err)
		}
	}()

	fmt.Printf("Starting listening gateway at %s\n", opts.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Errorf("Failed to listen and serve: %v\n", err)
		return err
	}
	return nil
}
