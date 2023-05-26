package main

import (
	"log"
	"mock_project/grpc/booking_grpc/handler"
	"mock_project/pb"
	"mock_project/grpc/booking_grpc/repository"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2228")
	if err != nil {
		panic(err)
	}

	// logger, _ := zap.NewProduction()
	// defer logger.Sync()

	s := grpc.NewServer()

	bookingRepository, err := repository.NewPostgresDB("host=localhost port=5432 user=root dbname=projectDB password=123 sslmode=disable")
	if err != nil {
		panic(err)
	}

	h, err := handler.NewBookingrHandler(bookingRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterBookingServiceRPCServer(s, h)

	log.Println("BOOKING server is listening at port 2228...")
	s.Serve(listen)
}
