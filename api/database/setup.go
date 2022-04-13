package database

import (
	"jjrepos/gonang/api/database/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Connect() {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{PrepareStmt: true, Logger: dbLogger})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Book{})
	Db = db
}
