package middleware

import (
	"booking-venue-api/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.JWT([]byte(config.LoadAuthConfig().AccessSecret))
}

// fungsi untuk men generate token
func CreateToken(id int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 200).Unix() //Token expires after 200 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.LoadAuthConfig().AccessSecret))
}

func ExtractToken(c echo.Context) (int, error) {
	loginToken := c.Get("user").(*jwt.Token)
	if loginToken.Valid {
		claims := loginToken.Claims.(jwt.MapClaims)
		id := int(claims["id"].(float64))
		return id, nil
	}
	return -1, fmt.Errorf("unauthorized")
}
