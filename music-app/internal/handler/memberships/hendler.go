package memberships

import (
	"music-app/internal/models/memberships"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=hendler.go -destination=handler_mock_test.go -package=memberships
type service interface {
	SignUp(req memberships.SignUpRequest) error
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
	route := h.Group("/memberships")
	route.POST("/signup", h.SignUp)
}
