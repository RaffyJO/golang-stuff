package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()

	pageSizeStr := c.Query("pageSize")
	pageIndexStr := c.Query("pageIndex")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("Page Size is invalid").Error(),
		})
	}

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("Page Index is invalid").Error(),
		})
	}

	response, err := h.postsSvc.GetAllPost(ctx, int64(pageSize), int64(pageIndex))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
