package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	v1 "github.com/namtx/grpc-ecommerce/pkg/api/v1"
	"google.golang.org/grpc"
)

func RunProductServiceServer(ctx context.Context, v1API v1.ProductServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterProductServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("Shutting down gRPC server")
			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Println("Starting gRPC server...")
	return server.Serve(listen)
}

func RunOrderServiceServer(ctx context.Context, v1API v1.OrderServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterOrderServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("Shuting down Order gRPC server...")
			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Println("Starting Order gRPC server...")
	return server.Serve(listen)
}
