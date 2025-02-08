package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	SecretKey                   string
	DatabaseUsername            string
	DatabasePassword            string
	DatabaseHost                string
	DatabasePort                string
	DatabaseName                string
	EnableDatabaseAutomigration bool
	CloudinaryFolder            string
	CloudinaryCloudName         string
	CloudinaryAPIKey            string
	CLoudinaryAPISecret         string
	JWTExpirationTime           int64
}

func LoadConfig() (config *Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRET_KEY")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	enableDatabaseAutomigration, _ := strconv.ParseBool(os.Getenv("ENABLE_DATABASE_AUTOMIGRATION"))
	cloudinaryFolder := os.Getenv("CLOUDINARY_FOLDER")
	cloudinaryCloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudinaryAPIKey := os.Getenv("CLOUDINARY_API_KEY")
	cLoudinaryAPISecret := os.Getenv("CLOUDINARY_API_SECRET")
	JWTExpirationTime := time.Now().Add(time.Second * 20).Unix()

	return &Config{
		SecretKey:                   secretKey,
		DatabaseUsername:            databaseUsername,
		DatabasePassword:            databasePassword,
		DatabaseHost:                databaseHost,
		DatabasePort:                databasePort,
		DatabaseName:                databaseName,
		EnableDatabaseAutomigration: enableDatabaseAutomigration,
		CloudinaryFolder:            cloudinaryFolder,
		CloudinaryCloudName:         cloudinaryCloudName,
		CloudinaryAPIKey:            cloudinaryAPIKey,
		CLoudinaryAPISecret:         cLoudinaryAPISecret,
		JWTExpirationTime:           JWTExpirationTime,
	}
}
