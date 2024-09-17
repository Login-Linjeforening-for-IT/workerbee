package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"gitlab.login.no/tekkom/web/beehive/admin-api/api"
	"gitlab.login.no/tekkom/web/beehive/admin-api/config"
	db "gitlab.login.no/tekkom/web/beehive/admin-api/db/sqlc"
	"gitlab.login.no/tekkom/web/beehive/admin-api/service"
	"gitlab.login.no/tekkom/web/beehive/admin-api/sessionstore"
	"gitlab.login.no/tekkom/web/beehive/admin-api/token"
)

type DBConfig struct {
	DBHost string `config:"DB_HOST" default:"localhost"`
	DBPort string `config:"DB_PORT" default:"5432"`
	DBUser string `config:"DB_USER" default:"root"`
	DBPass string `config:"DB_PASS" default:"secret"`
	DBName string `config:"DB_NAME" default:"beehivedb"`
}

type TLSConfig struct {
	Enabled bool   `config:"TLS_ENABLED" default:"true"`
	Cert    string `config:"TLS_CERT" default:"cert.pem"`
	Key     string `config:"TLS_KEY" default:"key.pem"`
}

type TokenConfig struct {
	AccessTokenSymmetricKey  string        `config:"TOKEN_ACCESS_TOKEN_SYMMETRIC_KEY"`
	RefreshTokenSymmetricKey string        `config:"TOKEN_REFRESH_TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration      time.Duration `config:"TOKEN_ACCESS_TOKEN_DURATION" default:"15m"`
	RefreshTokenDuration     time.Duration `config:"TOKEN_REFRESH_TOKEN_DURATION" default:"720h"`
}

type Oauth2Config struct {
	ClientID     	string `config:"OAUTH2_CLIENT_ID"`
	ClientSecret 	string `config:"OAUTH2_CLIENT_SECRET"`
	RedirectURL  	string `config:"OAUTH2_REDIRECT_URL"`
	RedirectClient	string `config:OAUTH2_FRONTEND_REDIRECT_URL`

	AuthentikBaseURL string `config:"OAUTH2_AUTHENTIK_BASE_URL" default:"https://authentik.login.no"`
}

var (
	configFile = flag.String("config", ".env", "path to config file")
)

func init() {
	flag.Parse()
}

func guard(err error) {
	if err != nil {
		panic(fmt.Errorf("%T %w", err, err))
	}
}

func main() {
	// Load config
	log.Println("Loading config...")
	fileopt := config.WithFile(*configFile)
	dbConf := config.MustLoad[DBConfig](fileopt)
	sessionStoreConf := config.MustLoad[sessionstore.RedisConfig](fileopt)
	apiConf := config.MustLoad[api.Config](fileopt)
	tlsConf := config.MustLoad[TLSConfig](fileopt)
	tokenConf := config.MustLoad[TokenConfig](fileopt)
	oauth2Conf := config.MustLoad[Oauth2Config](fileopt)

	// Auth and tokens
	log.Println("Creating token makers...")
	accessTokenMaker, err := token.NewPasetoMaker(tokenConf.AccessTokenSymmetricKey, token.AccessToken, tokenConf.AccessTokenDuration)
	guard(err)

	refreshTokenMaker, err := token.NewPasetoMaker(tokenConf.RefreshTokenSymmetricKey, token.RefreshToken, tokenConf.RefreshTokenDuration)
	guard(err)

	authentik := api.AuthentikOauth2Config(oauth2Conf.AuthentikBaseURL, oauth2Conf.ClientID, oauth2Conf.ClientSecret, oauth2Conf.RedirectURL, oauth2Conf.RedirectClient)

	// Connect to database
	log.Println("Connecting to database...")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConf.DBUser, dbConf.DBPass, dbConf.DBHost, dbConf.DBPort, dbConf.DBName)

	conn, err := sql.Open("postgres", dsn)
	guard(err)
	defer conn.Close()
	guard(conn.Ping())

	store := db.NewStore(conn)
	service := service.NewService(store)

	// Connect to session store
	log.Println("Connecting to session store...")
	client := redis.NewClient(&redis.Options{
		Addr:     sessionStoreConf.Addr,
		Username: sessionStoreConf.Username,
		Password: sessionStoreConf.Password,
		DB:       sessionStoreConf.DB,
	})
	guard(client.Ping(context.Background()).Err())

	sessionStore := sessionstore.New(client)

	// Start server
	log.Println("Starting server...")
	server := api.NewServer(apiConf,
		service, sessionStore,
		authentik,
		accessTokenMaker, 
		refreshTokenMaker,
	)

	if tlsConf.Enabled {
		guard(server.StartTLS(tlsConf.Cert, tlsConf.Key))
	} else {
		guard(server.Start())
	}
}
