package main

import (
	"log"
	"net"

	"github.com/KennedySurianto/MICROSERVICE_GPT/controller"
	"github.com/KennedySurianto/MICROSERVICE_GPT/repository"
	"github.com/KennedySurianto/MICROSERVICE_GPT/service"

	"google.golang.org/grpc"
)

func main() {
	repo := repository.NewOrderRepo()
	svc := service.NewOrderService(repo)
	ctrl := controller.NewOrderController(svc)

	lis, _ := net.Listen("tcp", ":50052")
	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, ctrl)

	log.Println("Order service running on :50052")
	s.Serve(lis)
}
