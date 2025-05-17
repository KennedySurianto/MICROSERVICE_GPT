package service

type User struct {
	ID    int64
	Name  string
	Email string
}

type UserRepository interface {
	GetUserByID(id int64) (*User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int64) (*User, error) {
	return s.repo.GetUserByID(id)
}
