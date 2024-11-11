package memberships

import (
	"context"
	"errors"
	"forum-site-app/internal/model/memberships"
	"forum-site-app/pkg/jwt"
	tokenUtils "forum-site-app/pkg/token"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipsRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Str("email", req.Email).Msg("Failed to get user")
	}
	if user == nil {
		return "", "", errors.New("Email not already exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", "", errors.New("Email or password is incorrect")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", "", err
	}

	existingRefreshToken, err := s.membershipsRepo.GetRefreshToken(ctx, user.ID, user.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token")
		return "", "", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUtils.GenerateRefreshToken()
	if refreshToken == "" {
		return "", "", errors.New("Failed to generate refresh token")
	}

	err = s.membershipsRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    user.UpdatedAt.Add(time.Hour * 24 * 7),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(user.ID, 10),
		UpdatedBy:    strconv.FormatInt(user.ID, 10),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert refresh token")
		return "", "", err
	}

	return token, refreshToken, nil
}
