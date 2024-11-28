package main

import (
	"fmt"
	"log"
	"microservice_grpc_cart/db"
	"microservice_grpc_cart/pb/cart"
	"microservice_grpc_cart/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := db.ConnectDatabse()
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	grpcServer := grpc.NewServer()
	cartService := &service.CartSeerviceServer{
		DB: db,
	}

	cart.RegisterCartSeerviceServer(grpcServer,cartService)
	lis,err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalf("Failed to listen on port 8080: %v",err)
	}

	fmt.Println("Server running on port :8080")
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to connect gRPC server %v",err)
	}
}
