package memberships

import (
	"music-app/internal/models/memberships"
	"music-app/internal/models/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req memberships.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.WebResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	err := h.service.SignUp(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.WebResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, response.WebResponse{
		Status:  "success",
		Message: "Successfully signed up",
		Data:    nil,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req memberships.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.WebResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	accessToken, err := h.service.Login(req)
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
		Message: "Successfully logged in",
		Data: memberships.LoginResponse{
			AccessToken: accessToken,
		},
	})
}
