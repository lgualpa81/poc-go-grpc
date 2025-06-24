package order

import (
	"fmt"

	"github.com/lgualpa81/go-grpc-api-gateway/pkg/config"
	pb "github.com/lgualpa81/go-grpc-api-gateway/pkg/order/pb/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.NewClient(c.OrderSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("could not connect: %w", err)
	}
	return pb.NewOrderServiceClient(cc)
}
