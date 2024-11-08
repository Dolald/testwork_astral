package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Dolald/testwork_astral/configs"
	"github.com/Dolald/testwork_astral/internal/repository"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.User
}

func NewAuthService(repo repository.User) *AuthService {
	return &AuthService{repo: repo}
}

func gereratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(configs.Salt)))
}

func (a *AuthService) GenerateToken(ctx context.Context, username, password string) (string, error) {
	userId, err := a.repo.GetUser(ctx, username, gereratePasswordHash(password))
	if err != nil {
		return "", fmt.Errorf("GetUser failed: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(configs.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId.Id,
	})

	return token.SignedString([]byte(configs.SigngingKey))
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(configs.SigngingKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("ParseWithClaims failed: %w", err)
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
