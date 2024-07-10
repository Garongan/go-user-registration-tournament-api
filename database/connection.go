package database

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"go-user-registration-tournament/config"
	"go-user-registration-tournament/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dbUser := config.Config("DB_USER")
	dbPassword := config.Config("DB_PASSWORD")
	dbHost := config.Config("DB_HOST")
	dbPort := config.Config("DB_PORT")
	dbName := config.Config("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Info(dbUser)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	DB = db
	err = db.AutoMigrate(&model.Account{}, &model.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
