package main

import (
	"database/sql"
	"fmt"

	"git.logntnu.no/tekkom/web/beehive/admin-api/api"
	"git.logntnu.no/tekkom/web/beehive/admin-api/config"
	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"git.logntnu.no/tekkom/web/beehive/admin-api/service"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	DBHost string `config:"DB_HOST" default:"localhost"`
	DBPort string `config:"DB_PORT" default:"5432"`
	DBUser string `config:"DB_USER" default:"root"`
	DBPass string `config:"DB_PASS" default:"secret"`
	DBName string `config:"DB_NAME" default:"beehivedb"`
}

func guard(err error) {
	if err != nil {
		panic(fmt.Errorf("%T %w", err, err))
	}
}

func main() {
	conf := config.MustLoad[DBConfig](config.WithFile(".env"))
	apiConf := config.MustLoad[api.Config](config.WithFile(".env"))

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)

	conn, err := sql.Open("postgres", dsn)
	guard(err)
	defer conn.Close()

	err = conn.Ping()
	guard(err)

	store := db.NewStore(conn)
	service := service.NewService(store)

	server := api.NewServer(apiConf, service)

	guard(server.Start())
}
