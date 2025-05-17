package repository

import "github.com/KennedySurianto/MICROSERVICE_GPT/service"

type InMemoryOrderRepo struct {
	lastID int64
	orders []*service.Order
}

func NewOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{}
}

func (r *InMemoryOrderRepo) Create(order *service.Order) (*service.Order, error) {
	r.lastID++
	order.ID = r.lastID
	r.orders = append(r.orders, order)
	return order, nil
}
