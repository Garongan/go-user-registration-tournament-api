package database

import (
	"go-user-registration-tournament/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	connectionString := "root:password@tcp(localhost:3306)/user_registration_tournament_db"

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	DB = db
	err = db.AutoMigrate(&models.Account{}, &models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
