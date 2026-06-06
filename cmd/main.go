package main

import (
	"detrox/internal/user"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"not null"`
	Email    string `json:"email" validate:"required" gorm:"uniqueIndex;not null"`
	Password string `json:"password" validate:"required,min=6" gorm:"not null"`
}

func main() {
	// Database connection
	dsn := "host=localhost user=manik password=14062021 dbname=pwd-manager port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&User{})
	fmt.Println("Database connected")

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RequestLogger())

	user.RegisterRoutes(e, db)

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Println("Server running on http://localhost:8080")
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
