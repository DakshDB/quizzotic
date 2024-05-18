package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// JWTMiddleware checks for a valid JWT token in the request header
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authorizationHeader := c.Request().Header.Get("Authorization")
        if authorizationHeader == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "authorization token required"})
        }

        tokenString := strings.TrimSpace(strings.TrimPrefix(authorizationHeader, "Bearer "))
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Ensure the signing method is what you expect
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }

            // Use the same secret as when you created the token
            return []byte(viper.GetString("JWT_SECRET")), nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
        }

        return next(c)
    }
}
