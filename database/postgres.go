package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DarkWorldCoder/auth/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var PRIVATE_KEY string

func ConnectToDB() {
	err := godotenv.Load()
	PRIVATE_KEY = "ayush"
	if err != nil {
		log.Fatal("Error loading env file \n", err)
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kathmandu",
		os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PORT"))

	log.Print("Connecting to Postgres Db...")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Field to connect to database. \n", err)
	}
	log.Println("Connected")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Print("Running the migrations ...")
	DB.AutoMigrate(&models.User{}, &models.Claims{})

}
