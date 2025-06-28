package auth

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

var secretKey []byte

func GenerateToken(userID int, username string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"iss":      "tasks-api",
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	})

	accessTokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return accessTokenString, nil
}

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
		}})
}
