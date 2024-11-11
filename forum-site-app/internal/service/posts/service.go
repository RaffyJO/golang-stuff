package posts

import (
	"context"
	"forum-site-app/internal/configs"
	"forum-site-app/internal/model/posts"
)

type postsRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error

	GetAllPost(ctx context.Context, limit, offset int64) (posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context, postID int64) (*posts.Post, error)

	GetCommentsByPostID(ctx context.Context, postID int64) ([]posts.Comment, error)
	CountLike(ctx context.Context, postID int64) (int64, error)
}

type service struct {
	cfg       *configs.Config
	postsRepo postsRepository
}

func NewService(cfg *configs.Config, postsRepo postsRepository) *service {
	return &service{
		cfg:       cfg,
		postsRepo: postsRepo,
	}
}
