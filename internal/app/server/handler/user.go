package handler

import (
	"fmt"
	"net/http"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/app/server/jwt"
	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

type userResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func newUserResponse(user *model.User) *userResponse {
	return &userResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}

type tokenResponse struct {
	AccessToken string       `json:"accessToken"`
	ExpiresIn   int64        `json:"expiresIn"`
	User        userResponse `json:"user"`
}

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
	return &tokenResponse{
		AccessToken: token,
		ExpiresIn:   86400,
		User: userResponse{
			ID:       u.ID,
			Username: u.Username,
		},
	}
}

func (h *Handler) Login(c *echo.Context) error {
	// redirectUri := c.QueryParam("redirect_uri")
	req := &userLoginRequest{}
	if code, err := req.bind(c); err != nil {
		return c.JSON(code, utils.NewError(err))
	}
	users, err := config.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	var token string
	for idx, user := range users {
		if req.Username == user.Username && req.Password == user.Password {
			token = newTokenResponse(&model.User{ID: uint(idx), Username: user.Username}).AccessToken
		}
	}
	if token == "" {
		return c.JSON(http.StatusForbidden, utils.NewError(fmt.Errorf("用户或密码错误")))
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
	}
	c.SetCookie(cookie)
	// if !strings.HasPrefix(redirectUri, "/h5") || !strings.HasPrefix(redirectUri, "/console") {
	// 	redirectUri = "/h5"
	// }
	// return c.Redirect(http.StatusFound, redirectUri)
	return c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) CurrentUser(c *echo.Context) error {
	claims := jwt.GetClaims(c)
	return c.JSON(http.StatusOK, newUserResponse(&model.User{ID: claims.UserID, Username: claims.Username}))
}
