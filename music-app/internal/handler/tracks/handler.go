package tracks

import (
	"context"
	"music-app/internal/middleware"
	"music-app/internal/models/spotify"
	"music-app/internal/models/track_activities"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=hendler.go -destination=handler_mock_test.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error)
	UpsertTrackActivity(userID uint, request track_activities.TrackActivityRequest) error
	GetRecommendations(userID uint, limit int, trackID string) (*spotify.RecommendationResponse, error)
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/tracks")
	route.Use(middleware.AuthMiddleware())
	route.GET("/search", h.Search)
	route.POST("/track-activity", h.UpsertTrackActivity)
	route.GET("/recommendations", h.GetRecommendations)
}
