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
	repo := repository.NewUserRepo()
	svc := service.NewUserService(repo)
	ctrl := controller.NewUserController(svc)

	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, ctrl)

	log.Println("User service running on :50051")
	s.Serve(lis)
}
