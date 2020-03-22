package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	v1 "github.com/namtx/grpc-ecommerce/pkg/api/v1"
	"google.golang.org/grpc"
)

const (
	apiVersion = "v1"
)

func main() {
	makeOrderServiceRequest()
}

func makeOrderServiceRequest() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	searchStream, err := c.SearchOrders(ctx, &wrappers.StringValue{Value: "Google"})

	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			break
		}

		log.Print("Search Result:", searchOrder)
	}

	orders := []*v1.Order{
		&v1.Order{Id: "001"},
		&v1.Order{Id: "002"},
		&v1.Order{Id: "003"},
	}

	updateStream, err := c.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", c, err)
	}

	for _, order := range orders {
		if err := updateStream.Send(order); err != nil {
			log.Fatalf("%v.Send(%v) = %v", updateStream, order, err)
		}
	}

	updateResponse, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRec() got error %v, want %v", updateStream, err, nil)
	}

	log.Printf("UpdateOrders response: %s", updateResponse)
}

func makeProductServiceRequest() {
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
