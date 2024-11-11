package memberships

import (
	"context"
	"errors"
	"forum-site-app/internal/model/memberships"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipsRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("Username or email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := &memberships.UserModel{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.membershipsRepo.CreateUser(ctx, model)
}
