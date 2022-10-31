package main

import (
	"fmt"
	"log"
	"net"
	"predict/config"
	"predict/predict"
	"predict/repository"
	"predict/usecase"

	"google.golang.org/grpc"
)

func main() {
	db := config.ConnectDB()
	PredictRepository := repository.InitPredictRepository(db)
	predictUsecase := usecase.InitUserUsecase(PredictRepository)

	s := predict.InitServer(predictUsecase)

	grpcServer := grpc.NewServer()

	predict.RegisterPredictServiceServer(grpcServer, &s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7000))
	if err != nil {
		log.Println("failed to listen: ", err)
	}
	fmt.Println("Listen to port 7000")

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("failed to serve: ", err)
	}
}
