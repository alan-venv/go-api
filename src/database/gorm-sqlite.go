package database

import (
	"example/go-api/src/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func clean() {
	err := os.Remove("sqlite.db")
	if err != nil {
		fmt.Println(err)
	}
}

func StartSqlite() *gorm.DB {
	clean() //! Just for SQLite

	connection, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	config, _ := connection.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return connection
}
