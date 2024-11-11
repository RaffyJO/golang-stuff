package memberships

import (
	"context"
	"forum-site-app/internal/middleware"
	"forum-site-app/internal/model/memberships"

	"github.com/gin-gonic/gin"
)

type membershipsService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipsSvc membershipsService
}

func NewHandler(api *gin.Engine, membershipsSvc membershipsService) *Handler {
	return &Handler{
		Engine:         api,
		membershipsSvc: membershipsSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)

	routeRefreshToken := h.Group("/memberships")
	routeRefreshToken.Use(middleware.AuthRefreshMiddleware())
	routeRefreshToken.POST("/refresh-token", h.RefreshToken)
}
