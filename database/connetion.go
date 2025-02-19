package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "nipun1234"
    dbname   = "squrepos"
)
var Db *sql.DB
var err error

func Connect(){
	ConnectString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",host, user, password, dbname, port)	
	Db, err = sql.Open("postgres", ConnectString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Println("Database connection established successfully")
}