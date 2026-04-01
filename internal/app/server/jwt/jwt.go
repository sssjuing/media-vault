package jwt

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

type JwtCustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var JWTSecret []byte

func init() {
	secretKey := config.GetConfig().GetString("server.secret_key")
	if secretKey == "" {
		log.Fatal("secret key is not set")
	}
	JWTSecret = []byte(secretKey)
}

func GenerateJwtToken(claims *JwtCustomClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString(JWTSecret)
	return t
}

func GetClaims(c *echo.Context) *JwtCustomClaims {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return nil
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims
	}
	return nil
}

func ParseJwtToken(tokenStr string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (any, error) {
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok {
		return claims, err
	}
	return nil, fmt.Errorf("fail to convert to JwtCustomClaims")
}
