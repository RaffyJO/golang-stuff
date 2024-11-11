package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle   string   `json:"PostTitle" binding:"required"`
		PostContent string   `json:"PostContent" binding:"required"`
		PostHastags []string `json:"PostHastags" binding:"required"`
	}
)

type (
	PostModel struct {
		ID          int64     `db:"id"`
		UserID      int64     `db:"user_id"`
		PostTitle   string    `db:"post_title"`
		PostContent string    `db:"post_content"`
		PostHastags string    `db:"post_hastags"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
		CreatedBy   string    `db:"created_by"`
		UpdatedBy   string    `db:"updated_by"`
	}
)

type (
	GetAllPostResponse struct {
		Data       []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID          int64    `json:"id"`
		UserID      int64    `json:"userId"`
		Username    string   `json:"username"`
		PostTitle   string   `json:"postTitle"`
		PostContent string   `json:"postContent"`
		PostHastags []string `json:"postHastags"`
		IsLiked     bool     `json:"isLiked"`
	}

	Pagination struct {
		Limit  int64 `json:"limit"`
		Offset int64 `json:"offset"`
	}

	GetPostResponse struct {
		PostDetail Post      `json:"postDetail"`
		LikeCount  int64     `json:"likeCount"`
		Comments   []Comment `json:"comments"`
	}

	Comment struct {
		ID             int64  `json:"id"`
		UserID         int64  `json:"userId"`
		Username       string `json:"username"`
		CommentContent string `json:"commentContent"`
	}
)
