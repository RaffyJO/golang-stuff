package tracks

import (
	"music-app/internal/models/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	ctx := c.Request.Context()

	query := c.Query("query")
	pageSizeStr := c.Query("pageSize")
	pageIndexStr := c.Query("pageIndex")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		pageIndex = 1
	}

	userID := c.GetUint("userID")
	responseData, err := h.service.Search(ctx, query, pageSize, pageIndex, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.WebResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.WebResponse{
		Status:  "success",
		Message: "Successfully get search results",
		Data:    &responseData,
	})
}
