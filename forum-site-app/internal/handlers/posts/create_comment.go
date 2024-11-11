package posts

import (
	"errors"
	"forum-site-app/internal/model/posts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.New("Invalid Post ID").Error(),
		})
		return
	}

	userID := c.GetInt64("userID")
	err = h.postsSvc.CreateComment(ctx, postID, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
