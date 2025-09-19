package db

import (
	"log"
	"workerbee/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init() *sqlx.DB {
	DB, err := sqlx.Connect("postgres", config.DB_url)
	if err != nil {
		log.Fatalln("Unable to connect to db: ", err)
	}
	log.Println("Connected to DB")
	return DB
}
