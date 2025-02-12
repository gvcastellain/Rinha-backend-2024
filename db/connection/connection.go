package connection

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

func NewDbConn() {
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_NAME"))

	
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		return
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		return
	}
}

func GetDB() *sql.DB {
	return db
}
