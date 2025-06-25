package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/lgualpa81/go-grpc-api-gateway/pkg/product/pb/gen"
)

type CreateProductRequestBody struct {
	Name  string  `json:"name"`
	Stock int64   `json:"stock"`
	Price float64 `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateProductRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Printf("Sending to gRPC - Name: %s, Stock: %d, Price: %f\n", body.Name, body.Stock, body.Price)
	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
