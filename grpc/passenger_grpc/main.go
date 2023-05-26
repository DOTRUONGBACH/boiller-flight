package main

import (
	"log"
	"mock_project/grpc/passenger_grpc/handler"
	"mock_project/pb"
	"mock_project/grpc/passenger_grpc/repository"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2226")
	if err != nil {
		panic(err)
	}

	// logger, _ := zap.NewProduction()
	// defer logger.Sync()

	s := grpc.NewServer()

	passengerRepository, err := repository.NewPostgresDB("host=localhost port=5432 user=root dbname=projectDB password=123 sslmode=disable")
	if err != nil {
		panic(err)
	}

	h, err := handler.NewPassengerrHandler(passengerRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterPassengerServiceRPCServer(s, h)

	log.Println("PASSENGER server is listening at port 2226...")
	s.Serve(listen)
}
