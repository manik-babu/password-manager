package config

import (
	"detrox/internal/user"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(env *Env) *gorm.DB {
	dsn := env.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&user.User{})
	fmt.Println("Database connected")
	return db

}
