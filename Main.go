package main

import (
	"Contact/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=toluak@28 dbname=Contact port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		fmt.Println("Successfully connected to the database!")
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate the database:", err)
	} else {
		fmt.Println("Database migrated successfully!")
	}
	err = db.AutoMigrate(&model.ContactUsers{})
	if err != nil {
		log.Fatal("Failed to migrate the database:", err)
	} else {
		fmt.Println("Database migrated successfully!")
	}
}
