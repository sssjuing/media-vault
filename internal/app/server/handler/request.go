package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
)

func validateRequest[T any](c *echo.Context, req *T) error {
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return nil
}

type userLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *userLoginRequest) bind(c *echo.Context) error {
	return validateRequest(c, r)
}
