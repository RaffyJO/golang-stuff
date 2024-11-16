package memberships

import (
	"errors"
	"music-app/internal/models/memberships"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignUp(req memberships.SignUpRequest) error {
	exitingUser, err := s.repository.GetUser(req.Email, req.Username, 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Failed to get user")
		return err
	}
	if exitingUser != nil {
		return errors.New("Email or username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Failed to hash password")
		return err
	}

	model := memberships.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.repository.CreateUser(model)
}
