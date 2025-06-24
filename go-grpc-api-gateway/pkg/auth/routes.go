package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lgualpa81/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/lgualpa81/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	authServiceClient, _ := InitServiceClient(c)
	svc := &ServiceClient{
		Client: authServiceClient,
	}
	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
