package service

type Order struct {
	ID       int64
	UserID   int64
	Item     string
	Quantity int32
}

type OrderRepository interface {
	Create(order *Order) (*Order, error)
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(r OrderRepository) *OrderService {
	return &OrderService{repo: r}
}

func (s *OrderService) CreateOrder(userID int64, item string, qty int32) (*Order, error) {
	return s.repo.Create(&Order{
		UserID:   userID,
		Item:     item,
		Quantity: qty,
	})
}
