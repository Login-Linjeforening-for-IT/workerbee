package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/config"
)

func Init() *sqlx.DB {
	DB, err := sqlx.Connect("postgres", config.DB_url)
	if err != nil {
		log.Fatalln("Unable to connect to db: ", err)
	}
	log.Println("Connected to DB")
	return DB
}
