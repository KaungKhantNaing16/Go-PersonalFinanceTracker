package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "finance_tracker"
)

var (
	currentDateTime = time.Now()
	CreateAt        = currentDateTime
	UpdatedAt       = currentDateTime
	Deletedat       = currentDateTime
)

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbName)
}

// NewDatabase creates a new database connection pool
func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to DB %s successfully\n", dbname)

	return db
}
