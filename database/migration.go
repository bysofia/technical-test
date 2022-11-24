package database

import (
	"fmt"
	"nutech/models"
	"nutech/pkg/postgre"
)

func RunMigration() {
	err := postgre.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}