package main

import (
	"context"
	"log"
	"time"

	v1 "github.com/namtx/grpc-ecommerce/pkg/api/v1"
	"google.golang.org/grpc"
)

const (
	apiVersion = "v1"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req1 := &v1.Product{
		Name:        "iPhone 11",
		Description: "All-new dual-camera system with Ultra Wide and Night mode. ",
	}
	res1, err := c.AddProduct(ctx, req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)
}
