package service

import (
	"context"
	"net/http"

	"github.com/lgualpa81/go-grpc-order-svc/pkg/client"
	"github.com/lgualpa81/go-grpc-order-svc/pkg/db"
	"github.com/lgualpa81/go-grpc-order-svc/pkg/models"
	pb "github.com/lgualpa81/go-grpc-order-svc/pkg/pb/gen"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	H          db.Handler
	ProductSvc client.ProductServiceClient
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := s.ProductSvc.FindOne(req.ProductId)
	//fmt.Printf("CreateOrder datos recibidos (raw): %+v\n", req)
	//fmt.Printf("CreateOrder product: %+v", product)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{Status: product.Status, Error: product.Error}, nil
	} else if *product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: "Stock too less"}, nil
	}

	order := models.Order{
		Price:     *product.Data.Price,
		ProductId: product.Data.Id,
		Quantity:  req.Quantity,
		UserId:    req.UserId,
	}

	s.H.DB.Create(&order)

	res, err := s.ProductSvc.DecreaseStock(req.ProductId, order.Id, req.Quantity)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
