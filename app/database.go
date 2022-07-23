package app

import (
	"database/sql"
	"log"
	"os"
	"rahadiangg/category-api/helper"
	"time"

	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Println("Can't find .env file")
	}

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+":3306)/"+db_name)
	helper.PanicIfError(err)

	log.Println("Connected to database")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
