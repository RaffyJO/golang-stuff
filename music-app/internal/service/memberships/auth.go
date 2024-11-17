package memberships

import (
	"errors"
	"music-app/internal/models/memberships"
	"music-app/pkg/jwt"

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

func (s *service) Login(req memberships.LoginRequest) (string, error) {
	userDetail, err := s.repository.GetUser(req.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Failed to get user from database")
		return "", err
	}

	if userDetail == nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	accessToken, err := jwt.CreateToken(int64(userDetail.ID), userDetail.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate token")
		return "", err
	}

	return accessToken, nil
}
