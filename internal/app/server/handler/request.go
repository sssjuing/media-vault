package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func validateRequest[T any](c *echo.Context, req *T) (int, error) {
	if err := c.Bind(req); err != nil {
		return http.StatusBadRequest, err
	}
	if err := c.Validate(req); err != nil {
		return http.StatusUnprocessableEntity, err
	}
	return 0, nil
}

type userLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *userLoginRequest) bind(c *echo.Context) (int, error) {
	return validateRequest(c, r)
}
