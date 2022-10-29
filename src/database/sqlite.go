package database

import (
	"example/go-api/src/models"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var connection *gorm.DB

func Get() *gorm.DB {
	return connection
}

func clean() {
	err := os.Remove("test.db")
	if err != nil {
		fmt.Println(err)
	}
}

func Start() {
	clean() //! Just for SQLite

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	connection = db

	config, _ := connection.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	connection.AutoMigrate(&models.User{})
}
