package main

import (
	"database/sql"
	"fmt"

	"git.logntnu.no/tekkom/web/beehive/admin-api/api"
	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"git.logntnu.no/tekkom/web/beehive/admin-api/service"
	_ "github.com/lib/pq"
)

type Config struct {
	DBHost string `config:"DB_HOST" default:"localhost"`
	DBPort string `config:"DB_PORT" default:"5432"`
	DBUser string `config:"DB_USER" default:"postgres"`
	DBPass string `config:"DB_PASS" default:"postgres"`
	DBName string `config:"DB_NAME" default:"postgres"`

	Port string `config:"PORT" default:"8080"`
}

func guard(err error) {
	if err != nil {
		panic(fmt.Errorf("%T %w", err, err))
	}
}

func main() {
	conf := Config{
		DBHost: "localhost",
		DBPort: "5432",
		DBUser: "root",
		DBPass: "secret",
		DBName: "beehivedb",
		Port:   "8080",
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)

	conn, err := sql.Open("postgres", dsn)
	guard(err)
	defer conn.Close()

	err = conn.Ping()
	guard(err)

	store := db.NewStore(conn)
	service := service.NewService(store)

	server := api.NewServer(&api.Config{
		Port: conf.Port,
	}, service)

	guard(server.Start())
}
