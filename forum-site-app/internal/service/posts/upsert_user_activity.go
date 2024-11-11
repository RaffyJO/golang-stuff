package posts

import (
	"context"
	"errors"
	"forum-site-app/internal/model/posts"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	userActivity, err := s.postsRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user activity")
		return err
	}

	if userActivity == nil {
		if !request.IsLiked {
			return errors.New("user activity not found")
		}
		err = s.postsRepo.CreateUserActivity(ctx, model)
	} else {
		err = s.postsRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("Failed to update or create user activity")
		return err
	}

	return nil
}
