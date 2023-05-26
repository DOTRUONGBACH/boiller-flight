package main

import (
	"log"
	"mock_project/grpc/ticket_grpc/handler"
	"mock_project/pb"
	"mock_project/grpc/ticket_grpc/repository"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2227")
	if err != nil {
		panic(err)
	}

	// logger, _ := zap.NewProduction()
	// defer logger.Sync()

	s := grpc.NewServer()

	ticketRepository, err := repository.NewPostgresDB("host=localhost port=5432 user=root dbname=projectDB password=123 sslmode=disable")
	if err != nil {
		panic(err)
	}

	h, err := handler.NewTicketrHandler(ticketRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterTicketServiceRPCServer(s, h)

	log.Println("Ticket server is listening at port 2227...")
	s.Serve(listen)
}
