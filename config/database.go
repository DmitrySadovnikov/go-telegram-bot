package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB

func init() {
	godotenv.Load()
}

func OpenTestGORMConnection() {
	DB, _ = gorm.Open("testdb", "parseTime=True")
}

func OpenGORMConnection() {
	var err error
	connectString := fmt.Sprintf("%v dbname=%v binary_parameters=%v", postgresConnectString(), os.Getenv("POSTGRES_DNAME"), "yes")
	DB, err = gorm.Open("postgres", connectString)

	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("APP_ENV") == "production" || os.Getenv("APP_ENV") == "staging" {
		DB.LogMode(true)
	}

	return
}

func postgresConnectString() string {
	return fmt.Sprintf("host=%v user=%v port=%v sslmode=%v password=%v extra_float_digits=3",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"), os.Getenv("POSTGRES_PASSWORD"),
	)
}

func OpenSQLConnection() (db *sql.DB, err error) {
	return sql.Open("postgres", postgresConnectString())
}
