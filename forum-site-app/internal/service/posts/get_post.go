package posts

import (
	"context"
	"forum-site-app/internal/model/posts"

	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postsRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get post detail")
		return nil, err
	}

	comments, err := s.postsRepo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get comments")
		return nil, err
	}

	likeCount, err := s.postsRepo.CountLike(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get like count")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: *postDetail,
		LikeCount:  likeCount,
		Comments:   comments,
	}, nil
}
