package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/internal/config"
)

func Init() {
	db, err := sqlx.Connect("postgres", config.DB_url)
	if err != nil {
		log.Fatalln("Unable to connect to db: ", err)
	}
	defer db.Close()
	log.Println("Connected to DB")
}
