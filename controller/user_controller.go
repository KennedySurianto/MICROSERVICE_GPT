package controller

import (
	"context"

	"github.com/KennedySurianto/MICROSERVICE_GPT/service"
	"github.com/KennedySurianto/MICROSERVICE_GPT/proto"
)

type UserController struct {
	userService *service.UserService
	proto.UnimplementedUserServiceServer
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{userService: svc}
}

func (c *UserController) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, _ := c.userService.GetUser(req.Id)
	return &proto.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
