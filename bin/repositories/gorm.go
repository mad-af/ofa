package repositories

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgre() *gorm.DB {
	var connection = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"127.0.0.1",
		"postgres",
		"postgres",
		"coba",
		"5432",
		"disable",
	)

	var db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database postgre")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to create pool connection database postgre")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
