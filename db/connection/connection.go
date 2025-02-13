package connection

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "gab"
	dbname   = "rinhadb"
)

func NewDbConn() {
	createDB()

	var err error
	db, err = sql.Open("postgres", getPsql(dbname))

	if err != nil {
		return
	}

	err = db.Ping()

	if err != nil {
		return
	}

	sqlFile := "db\\sql\\scripts\\schemas.sql"
	sqlContent, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return
	}

	_, err = db.Exec(string(sqlContent))
	if err != nil {
		return
	}
}

func GetDB() *sql.DB {
	return db
}

func createDB() {
	newDB, err := sql.Open("postgres", getPsql("postgres"))
	if err != nil {
		return
	}

	defer newDB.Close()

	_, err = newDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbname))
	if err != nil {
		return
	}
	newDB.Ping()
}

func getPsql(dbname string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
