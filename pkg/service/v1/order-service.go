package v1

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	v1 "github.com/namtx/grpc-ecommerce/pkg/api/v1"
)

type orderServiceServer struct {
	orderMap map[string]*v1.Order
}

func NewOrderServiceServer() v1.OrderServiceServer {
	orders := map[string]*v1.Order{"001": &v1.Order{Items: []string{"Google"}}}

	return &orderServiceServer{orderMap: orders}
}

func (s *orderServiceServer) SearchOrders(searchQuery *wrappers.StringValue, stream v1.OrderService_SearchOrdersServer) error {
	for key, order := range s.orderMap {
		log.Print(key, order)

		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf("error sending to stream: ", err)
				}

				log.Print("Matching Order Found: " + key)
				break
			}
		}
	}

	return nil
}

func (s *orderServiceServer) UpdateOrders(stream v1.OrderService_UpdateOrdersServer) error {
	ordersStr := "Updated Order IDs: "

	for {
		order, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&wrappers.StringValue{Value: "Orders processed " + ordersStr})
		}

		s.orderMap[order.Id] = order

		log.Printf("Order ID %s: Updated", order.Id)
		ordersStr += order.Id + ", "
	}
}
