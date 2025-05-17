package repository

import "github.com/KennedySurianto/MICROSERVICE_GPT/service"

type InMemoryUserRepo struct {
	data map[int64]*service.User
}

func NewUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		data: map[int64]*service.User{
			1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
		},
	}
}

func (r *InMemoryUserRepo) GetUserByID(id int64) (*service.User, error) {
	return r.data[id], nil
}
