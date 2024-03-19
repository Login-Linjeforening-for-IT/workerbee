package main

import (
	"database/sql"
	"flag"
	"fmt"

	"git.logntnu.no/tekkom/web/beehive/admin-api/api"
	"git.logntnu.no/tekkom/web/beehive/admin-api/config"
	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"git.logntnu.no/tekkom/web/beehive/admin-api/image"
	"git.logntnu.no/tekkom/web/beehive/admin-api/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	DBHost string `config:"DB_HOST" default:"localhost"`
	DBPort string `config:"DB_PORT" default:"5432"`
	DBUser string `config:"DB_USER" default:"root"`
	DBPass string `config:"DB_PASS" default:"secret"`
	DBName string `config:"DB_NAME" default:"beehivedb"`
}

type DOConfig struct {
	DOKey     string `config:"DO_ACCESS_KEY_ID"`
	DOSecret  string `config:"DO_SECRET_ACCESS_KEY"`
	DORegion  string `config:"DO_REGION" default:"ams3"`
	DOBaseURL string `config:"DO_BASE_URL" default:"https://ams3.digitaloceanspaces.com"`
}

func guard(err error) {
	if err != nil {
		panic(fmt.Errorf("%T %w", err, err))
	}
}

var (
	configFile = flag.String("config", ".env", "path to config file")
)

func init() {
	flag.Parse()
}

func main() {
	conf := config.MustLoad[DBConfig](config.WithFile(*configFile))
	apiConf := config.MustLoad[api.Config](config.WithFile(*configFile))
	doConf := config.MustLoad[DOConfig](config.WithFile(*configFile))

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)

	conn, err := sql.Open("postgres", dsn)
	guard(err)
	defer conn.Close()

	err = conn.Ping()
	guard(err)

	store := db.NewStore(conn)
	service := service.NewService(store)

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(doConf.DOKey, doConf.DOSecret, ""),
		Endpoint:         aws.String(doConf.DOBaseURL),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String(doConf.DORegion),
	}

	doStore := image.NewDOStore(s3Config)

	server := api.NewServer(apiConf, service, doStore)

	guard(server.Start())
}
