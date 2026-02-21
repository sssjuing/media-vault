package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/app/server/jwt"
	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

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
			token = newTokenResponse(&model.User{ID: uint(idx), Username: user.Username}).Data.Token
		}
	}
	if token == "" {
		return c.JSON(http.StatusForbidden, utils.NewError(fmt.Errorf("用户不存在")))
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
