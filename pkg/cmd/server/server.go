package cmd

import (
	"context"

	"github.com/namtx/grpc-ecommerce/pkg/protocol/grpc"
	v1 "github.com/namtx/grpc-ecommerce/pkg/service/v1"
)

func RunProductServiceServer() error {
	ctx := context.Background()

	v1API := v1.NewProductServiceServer()
	port := "50051"

	return grpc.RunProductServiceServer(ctx, v1API, port)
}

func RunOrderServiceServer() error {
	ctx := context.Background()
	v1API := v1.NewOrderServiceServer()

	port := "50052"

	return grpc.RunOrderServiceServer(ctx, v1API, port)
}
