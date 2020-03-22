package v1

import (
	"context"

	"github.com/gofrs/uuid"
	v1 "github.com/namtx/grpc-ecommerce/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productServiceServer struct {
	productMap map[string]*v1.Product
}

func NewProductServiceServer() v1.ProductServiceServer {
	return &productServiceServer{}
}

func (s *productServiceServer) AddProduct(ctx context.Context, in *v1.Product) (*v1.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()

	if s.productMap == nil {
		s.productMap = make(map[string]*v1.Product)
	}

	s.productMap[in.Id] = in

	return &v1.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *productServiceServer) GetProduct(ctx context.Context, in *v1.ProductID) (*v1.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}
