package middleware

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS256.Name,
		SigningKey: []byte(config.SECRET_JWT),
	})
}

func CreateToken(id int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractToken(c echo.Context) (int) {
	headerData := c.Request().Header.Get("Authorization")
	dataAuth := strings.Split(headerData, " ")
	token := dataAuth[len(dataAuth)-1]
	datajwt, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_JWT), nil
	})

	if datajwt.Valid {
		claims := datajwt.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		return int(id)
	}

	return -1
}