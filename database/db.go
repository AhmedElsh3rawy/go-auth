package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/AhmedElsh3rawy/go-auth/users"
	_ "github.com/lib/pq"
)

var Query *users.Queries

func connectDB() *sql.DB {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr+"?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDB() {
	db := connectDB()
	Query = users.New(db)
}
