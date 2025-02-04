package config

import (
	"fmt"
	"log"
	"reservation/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() *gorm.DB {
	host := LoadConfig().DatabaseHost
	port := LoadConfig().DatabasePort
	user := LoadConfig().DatabaseUsername
	password := LoadConfig().DatabasePassword
	name := LoadConfig().DatabaseName

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, name, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if LoadConfig().EnableDatabaseAutomigration {
		go RunAutoMigration()
	}

	log.Printf("Connected to database: %s", name)

	return DB
}

func RunAutoMigration() {
	err := DB.AutoMigrate(
		&models.Users{},
		&models.Carousels{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
