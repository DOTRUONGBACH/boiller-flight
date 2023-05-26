package main

import (
	"log"
	"mock_project/grpc/account_grpc/handler"
	"mock_project/pb"
	"mock_project/grpc/account_grpc/repository"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2224")
	if err != nil {
		panic(err)
	}

	// logger, _ := zap.NewProduction()
	// defer logger.Sync()

	s := grpc.NewServer()

	accountRepository, err := repository.NewPostgresDB("host=localhost port=5432 user=root dbname=projectDB password=123 sslmode=disable")
	if err != nil {
		panic(err)
	}

	// defer userRepository.Close()

	h, err := handler.NewAccountrHandler(accountRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterAccountServiceRPCServer(s, h)

	// logger.Info("Listen at port: 2223")

	log.Println("ACCOUNT server is listening at port 2224...")
	s.Serve(listen)
}
