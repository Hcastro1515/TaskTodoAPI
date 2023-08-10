package db

import (
	"log"

	"github.com/hcastro1515/todosapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "postgres://pg:pass@localhost:5432/todos"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&models.Todo{})

	return db
}
