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
	SmtpHost                    string
	SmtpSenderName              string
	SmtpUsername                string
	SmtpPassword                string
	SmtpPort                    int
}

func LoadConfig() (config *Config) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
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
	JWTExpirationTime := time.Now().Add(time.Hour * 24).Unix()
	smtpHost := os.Getenv("SMTP_HOST")
	smtpSenderName := os.Getenv("SMTP_SENDER_NAME")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

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
		SmtpHost:                    smtpHost,
		SmtpSenderName:              smtpSenderName,
		SmtpUsername:                smtpUsername,
		SmtpPassword:                smtpPassword,
		SmtpPort:                    smtpPort,
	}
}
