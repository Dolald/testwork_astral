package handler

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/Dolald/testwork_astral/configs"
	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentify(c *gin.Context) {
	header := c.GetHeader(configs.AuthorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(configs.UserCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(configs.UserCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is not found")
		return 0, errors.New("user id is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is not found")
	}

	return idInt, nil
}

func validateLogin(login string) error {
	if len(login) < 8 {
		return errors.New("login must be at least 8 characters long")
	}
	if matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", login); !matched {
		return errors.New("login must contain only letters and digits")
	}
	return nil
}

func validatePassword(pswd string) error {
	if len(pswd) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, char := range pswd {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case !((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')):
			hasSpecial = true
		}
	}

	if !(hasUpper && hasLower && hasDigit && hasSpecial) {
		return errors.New("password must contain at least 2 letters in different cases, 1 digit, and 1 special character")
	}

	return nil
}
