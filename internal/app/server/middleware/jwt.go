package middleware

import (
	gojwt "github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/app/server/jwt"
)

func JWT() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		TokenLookup: "cookie:token",
		NewClaimsFunc: func(c *echo.Context) gojwt.Claims {
			return new(jwt.JwtCustomClaims)
		},
		SigningKey: jwt.JWTSecret,
	})
}
