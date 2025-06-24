package product

import (
	"fmt"

	"github.com/lgualpa81/go-grpc-api-gateway/pkg/config"
	pb "github.com/lgualpa81/go-grpc-api-gateway/pkg/product/pb/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.NewClient(c.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("could not connect:", err)
	}
	return pb.NewProductServiceClient(cc)
}
