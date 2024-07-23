package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Database struct {
	db *sql.DB
}

// constructor
func NewDatabase() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mariadbUser := os.Getenv("MARIADB_USER")
	mariadbPW := os.Getenv("MARIADB_PW")
	mariadbIP := os.Getenv("MARIADB_IP")
	mariadbConnDB := os.Getenv("MARIADB_CONNDB")
	additionalArgs := "?parseTime=true"
	strConn := mariadbUser + ":" + mariadbPW + mariadbIP + mariadbConnDB + additionalArgs

	// sql.open() is not a connection. When you use sql.Open() you get a handle for a database.
	// The database/sql package manages a pool of connections in the background, and doesn't open any connections until you need them.
	db, err := sql.Open("mysql", strConn)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// return the pointer to that struct
	// the `&` operator takes the address of this newly created struct.
	return &Database{db: db}, nil
}

func (db *Database) Close() {
	db.db.Close()
}

func (db *Database) GetDB() *sql.DB {
	return db.db
}
