package db

import (
	"errors"
	"fmt"
	"os"

	"crispypod.com/crispypod-backend/helpers"
	"crispypod.com/crispypod-backend/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbName = os.Getenv("DB_NAME")
	var dbUser = os.Getenv("DB_USER")
	var dbPassword = os.Getenv("DB_PASSWORD")
	if len(dbPort) == 0 {
		dbPort = "5432"
	}
	var dbConnString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
	}

	err = db.AutoMigrate(&models.DbUser{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create User table: %s", err.Error()))
	}

	if err = db.First(&models.DbUser{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		password, _ := helpers.HashPassword("crispy.pod")
		admin := models.DbUser{UserName: "admin", Password: password, ID: uuid.New(), IsAdmin: true, DisplayName: "Admin", Email: "admin@crispypod.com"}
		db.Create(&admin)
	}

	err = db.AutoMigrate(&models.Episode{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create Episode table: %s", err.Error()))
	}

	err = db.AutoMigrate(&models.SiteConfig{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create SiteConfig table: %s", err.Error()))
	}

	err = db.AutoMigrate(&models.DeployLog{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create DeployLog table: %s", err.Error()))
	}

	err = db.AutoMigrate(&models.Hook{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create Hooks table: %s", err.Error()))
	}

	err = db.AutoMigrate(&models.HookLog{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create HooksLog table: %s", err.Error()))
	}

	if err = db.First(&models.SiteConfig{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		siteConfig := models.SiteConfig{
			ID:              uuid.New(),
			SiteName:        "CrispyPod",
			SiteDescription: "Super awesome podcast!",
			SiteUrl:         "https://crispypod.com",
			SetupComplete:   false,
		}
		db.Create(&siteConfig)
	}

	DB = db
}
