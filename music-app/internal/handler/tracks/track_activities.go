package tracks

import (
	"music-app/internal/models/response"
	"music-app/internal/models/track_activities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertTrackActivity(c *gin.Context) {
	var req track_activities.TrackActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.WebResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	userID := c.GetUint("userID")
	err := h.service.UpsertTrackActivity(userID, req)
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
		Message: "Successfully upserted track activity",
		Data:    nil,
	})
}
