package handler

import (
	"net/http"

	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		h.logger.Errorf("BindJSON failed: %w", err)
		return
	}

	if err := validateLogin(input.Login); err != nil {
		h.logger.Errorf("validateLogin failed: %w", err)
		return
	}

	if err := validatePassword(input.Password); err != nil {
		h.logger.Errorf("validatePassword failed: %w", err)
		return
	}

	id, err := h.service.User.CreateUser(input)
	if err != nil {
		h.logger.Errorf("CreateUser failed: %w", err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		h.logger.Errorf("BindJSON failed: %w", err)
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		h.logger.Errorf("GenerateToken failed: %w", err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})

}
