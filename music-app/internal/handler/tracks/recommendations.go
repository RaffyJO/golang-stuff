package tracks

import (
	"music-app/internal/models/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRecommendations(c *gin.Context) {
	userID := c.GetUint("userID")
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	trackID := c.Query("trackID")
	if trackID == "" {
		c.JSON(http.StatusBadRequest, response.WebResponse{
			Status:  "error",
			Message: "trackID parameter is required",
			Data:    nil,
		})
		return
	}

	responseData, err := h.service.GetRecommendations(userID, limit, trackID)
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
		Message: "Successfully get recommendations",
		Data:    responseData,
	})
}
