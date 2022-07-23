package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var webPort = "80"
var counts int

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting Auth Service")
	conn := connectToDB()
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres is not ready...")
			counts++
		} else {
			log.Println("Postgres is ready...")
			return conn
		}
		if counts > 10 {
			log.Println("Postgress connection count exceeds...")
			return nil
		}
		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
	}
}
