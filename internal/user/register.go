package user

import (
	"detrox/internal/auth"
	"detrox/internal/middleware"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	repo := NewRepository(db)
	jwtService := auth.NewJWTService("jemnsokfgebi3492bd402b")
	service := NewService(repo, jwtService)
	handler := NewHandler(service)

	api := e.Group("/api/v1")
	api.POST("/auth/signup", handler.CreateUser)
	api.POST("/auth/login", handler.LoginUser)
	api.GET("/me", handler.GetMe, middleware.AuthMiddleWare(jwtService))
}
