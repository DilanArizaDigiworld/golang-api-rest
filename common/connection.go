package common

import (
	"golang-rest-api/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func GetConnection() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envStr_db_host := os.Getenv("DB_HOST")
	envStr_db_port := os.Getenv("DB_PORT")
	envStr_db_user := os.Getenv("DB_USER")
	envStr_db_pass := os.Getenv("DB_PASS")
	envStr_db_name := os.Getenv("DB_NAME")

	connStr := envStr_db_user + ":" + envStr_db_pass + "@tcp(" + envStr_db_host + ":" + envStr_db_port + ")/" + envStr_db_name + "?parseTime=true"

	db, error := gorm.Open("mysql", connStr)

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Print("Migrando...")

	db.AutoMigrate(&models.Persona{})
}
