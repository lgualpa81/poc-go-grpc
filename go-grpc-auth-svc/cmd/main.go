package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lgualpa81/go-grpc-auth-svc/pkg/config"
	"github.com/lgualpa81/go-grpc-auth-svc/pkg/db"
	pb "github.com/lgualpa81/go-grpc-auth-svc/pkg/pb/gen"
	"github.com/lgualpa81/go-grpc-auth-svc/pkg/services"
	"github.com/lgualpa81/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h, err := db.Init(c.DBUrl)
	if err != nil {
		log.Fatalln("Failed db config", err)
	}

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: c.TokenExpirationHours,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
