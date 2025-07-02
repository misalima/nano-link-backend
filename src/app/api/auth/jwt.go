package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

var secretKey []byte

func GenerateToken(userID uuid.UUID, username string) (string, error) {
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

func GetUserIDFromToken(c echo.Context) (uuid.UUID, error) {
	userVal := c.Get("user")
	if userVal == nil {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "No token found in context")
	}
	user, ok := userVal.(*jwt.Token)
	if !ok {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token type in context")
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims in token")
	}
	subVal, ok := claims["sub"]
	if !ok {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "No sub claim in token")
	}
	sub, ok := subVal.(string)
	if !ok {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "sub claim is not a string")
	}
	userID, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}
