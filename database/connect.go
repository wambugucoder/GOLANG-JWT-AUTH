package database

import (
	"fmt"
	"golang_auth/config"
	"golang_auth/models"
	"strconv"

	"github.com/jinzhu/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("PSQL_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("PSQL_HOST"), port, config.Config("PSQL_USER"), config.Config("PSQL_PASSWORD"), config.Config("PSQL_DBNAME")))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database has been Migrated")
}
