package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nsavelyeva/go-shopping/models"
	"log"
)

var DB *gorm.DB

func Setup() {
	db, err := gorm.Open("sqlite3", "items.db")
	if err != nil {
		log.Fatal("Failed to connect to the database!")
	}

	db.LogMode(false)
	db.AutoMigrate(&models.Item{})

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
