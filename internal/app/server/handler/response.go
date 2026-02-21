package handler

import (
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/sssjuing/media-vault/internal/app/server/jwt"
	"github.com/sssjuing/media-vault/internal/app/server/model"
)

type Response[T any] struct {
	Data    T      `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type tokenResponse = Response[struct {
	Token string `json:"token"`
}]

func newTokenResponse(u *model.User) *tokenResponse {
	claims := &jwt.JwtCustomClaims{
		UserID:   u.ID,
		Username: u.Username,
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  gojwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.GenerateJwtToken(claims)
	r := new(tokenResponse)
	r.Data.Token = token
	return r
}

type userResponse = Response[struct {
	*model.User
}]

func newUserResponse(user *model.User) *userResponse {
	r := new(userResponse)
	r.Data.User = user
	return r
}
