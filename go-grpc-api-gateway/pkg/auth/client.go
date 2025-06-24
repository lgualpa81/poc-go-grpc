package auth

import (
	"fmt"

	pb "github.com/lgualpa81/go-grpc-api-gateway/pkg/auth/pb/gen"
	"github.com/lgualpa81/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) (pb.AuthServiceClient, error) {
	cc, err := grpc.NewClient(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not connect: %w", err)
	}
	return pb.NewAuthServiceClient(cc), nil
}
