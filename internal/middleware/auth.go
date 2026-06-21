package middleware

import (
	"detrox/internal/auth"
	"detrox/internal/httpresponse"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

func AuthMiddleWare(jwtService auth.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			fmt.Println(authHeader)
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Ok:      false,
					Code:    401,
					Message: "No token provided",
				})
			}
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(401, httpresponse.Error{
					Ok:      false,
					Code:    401,
					Message: "No token provided",
				})
			}
			tokenStr := parts[1]
			fmt.Println("Token: ", tokenStr)
			claims, err := jwtService.ValidateToken(tokenStr)

			if err != nil {
				return c.JSON(401, httpresponse.Error{
					Ok:      false,
					Code:    401,
					Message: err.Error(),
				})
			}
			c.Set("user", claims)

			return next(c)
		}
	}
}
