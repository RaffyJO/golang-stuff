package memberships

import (
	"context"
	"forum-site-app/internal/configs"
	"forum-site-app/internal/model/memberships"
	"time"
)

type membershipsRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model *memberships.UserModel) error

	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
}

type service struct {
	cfg             *configs.Config
	membershipsRepo membershipsRepository
}

func NewService(cfg *configs.Config, membershipsRepo membershipsRepository) *service {
	return &service{
		cfg:             cfg,
		membershipsRepo: membershipsRepo,
	}
}
