package main

import (
	"log"
	"mock_project/grpc/customer_grpc/handler"
	"mock_project/grpc/customer_grpc/repository"
	"mock_project/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}

	// logger, _ := zap.NewProduction()
	// defer logger.Sync()

	s := grpc.NewServer()

	customerRepository, err := repository.NewPostgresDB("host=localhost port=5432 user=root dbname=projectDB password=123 sslmode=disable")
	if err != nil {
		panic(err)
	}

	// defer userRepository.Close()

	h, err := handler.NewCustomerrHandler(customerRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterCustomerServiceRPCServer(s, h)

	// logger.Info("Listen at port: 2223")

	log.Println("CUSTOMER server is listening at port 2223...")
	s.Serve(listen)
}
