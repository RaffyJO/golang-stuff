package posts

import (
	"context"
	"forum-site-app/internal/model/posts"

	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int64) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)
	response, err := s.postsRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all post")
		return response, err
	}
	return response, nil
}
