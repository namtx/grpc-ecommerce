package cmd

import (
	"context"

	"github.com/namtx/grpc-ecommerce/pkg/protocol/grpc"
	v1 "github.com/namtx/grpc-ecommerce/pkg/service/v1"
)

func RunServer() error {
	ctx := context.Background()

	v1API := v1.NewProductServiceServer()

	port := "50051"

	return grpc.RunServer(ctx, v1API, port)
}
