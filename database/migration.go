package database

import (
	"nutech/models"
	"nutech/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)

	if err != nil {
		
	}
}