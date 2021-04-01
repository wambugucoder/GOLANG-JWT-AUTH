package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

//DB INSTANCE
var db *gorm.DB

func database() {
	envfiles := godotenv.Load()

	if envfiles != nil {
		log.Fatal("Error Loading env files \n", envfiles)
	}
	//CREATE DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PSQL_HOSTNAME"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DBNAME"))

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connecting t db \n", err)
		os.Exit(2)
	}
	log.Print("Successfully connected to db")

}
