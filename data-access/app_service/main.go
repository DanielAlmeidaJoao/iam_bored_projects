package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	envErr := godotenv.Load("properties.env")

	if envErr != nil {
		log.Fatal(envErr)
	}
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_ADDRESS"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}

	// Get a databse handle
	fmt.Println(cfg.FormatDSN())
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("connected!")
}
