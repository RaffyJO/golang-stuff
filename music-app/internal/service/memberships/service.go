package memberships

import (
	"music-app/internal/configs"
	"music-app/internal/models/memberships"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=memberships
type repository interface {
	CreateUser(model memberships.User) error
	GetUser(email, username string, id uint) (*memberships.User, error)
}

type service struct {
	cfg        *configs.Config
	repository repository
}

func NewService(cfg *configs.Config, repository repository) *service {
	return &service{
		cfg:        cfg,
		repository: repository,
	}
}