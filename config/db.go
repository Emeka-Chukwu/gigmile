package config

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var counts int64

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

func ConnectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	log.Println(dsn, "check")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgress not yet ready ...")
			log.Println(err)
			log.Println(dsn)
			counts++
		} else {
			log.Println("Connceted to progress")
			return connection
		}

		if counts > 10 {
			log.Panic(err)
			return nil
		}
		log.Println("Backing off for two seconds")
		time.Sleep(2 * time.Second)
	}
}
