package database

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Need mysql driver
)

type databaseConnection struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

var dbFile = "database/dbcreds.json"

var dbInfo databaseConnection

// DB is a global variable for the application.
// This is ok in go...
var DB *sql.DB

func init() {
	loadDatabaseCredentialFromFile()

	DB = connection()
}

// Connection to the database and return a *sql.DB object.
func connection() *sql.DB {
	dbLogin := dbInfo.User + ":" + dbInfo.Password + "@tcp(" + dbInfo.Host + ")/" + dbInfo.Database

	db, err := sql.Open("mysql", dbLogin)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func loadDatabaseCredentialFromFile() {
	jsonFile, err := os.Open(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(data, &dbInfo)
}
