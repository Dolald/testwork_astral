package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Dolald/testwork_astral/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	id, err := h.service.User.CreateUser(ctx, input)
	if err != nil {
		h.logger.Errorf("CreateUser failed: %w", err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		h.logger.Errorf("BindJSON failed: %w", err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	token, err := h.service.Authorization.GenerateToken(ctx, input.Login, input.Password)
	if err != nil {
		h.logger.Errorf("GenerateToken failed: %w", err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}
