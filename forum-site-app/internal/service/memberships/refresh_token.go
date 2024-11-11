package memberships

import (
	"context"
	"errors"
	"forum-site-app/internal/model/memberships"
	"forum-site-app/pkg/jwt"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipsRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("Refresh token has expired")
	}

	if existingRefreshToken.RefreshToken != request.RefreshToken {
		return "", errors.New("Invalid refresh token")
	}

	user, err := s.membershipsRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
	}
	if user == nil {
		return "", errors.New("Email not already exists")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}

	return token, nil
}
