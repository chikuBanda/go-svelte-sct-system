package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sct-system/models"
)

var DB *gorm.DB

func StartDbConnection() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	println("Connected to database")

	err = db.AutoMigrate(
		&models.Beneficiary{},
		&models.Transfer{},
		&models.TransactionHistory{},
	)
	if err != nil {
		panic("failed to auto migrate")
	}
	println("Database migrated")

	DB = db
}
