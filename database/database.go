package database

import (
	"fmt"
	"os"

	"github.com/FabienDeborde/noas_projects/utils/logger"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DBConn *gorm.DB

func Init() {
	_, slogger := logger.Init()
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", dbHost, dbPort, dbName, dbUsername, dbPassword)

	conn, err := gorm.Open("postgres", dbURI)
	// database.DBConn, err = gorm.Open("sqlite3", "projects.db")
	if err != nil {
		slogger.Error(err)
		panic("Failed to connect to database")
	}
	DBConn = conn
	slogger.Infow("Database connection successfully opened!")

	slogger.Infow("Database migrated!")
}
