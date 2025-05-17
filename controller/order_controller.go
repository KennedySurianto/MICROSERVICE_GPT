package controller

import (
	"context"

	"github.com/KennedySurianto/MICROSERVICE_GPT/service"
	"github.com/KennedySurianto/MICROSERVICE_GPT/proto"
)

type OrderController struct {
	orderService *service.OrderService
	proto.UnimplementedOrderServiceServer
}

func NewOrderController(s *service.OrderService) *OrderController {
	return &OrderController{orderService: s}
}

func (c *OrderController) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	order, _ := c.orderService.CreateOrder(req.UserId, req.Item, req.Quantity)
	return &proto.CreateOrderResponse{
		OrderId: order.ID,
		Status:  "created",
	}, nil
}
