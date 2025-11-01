package db

import (
	"log"
	"workerbee/config"
	"workerbee/internal"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

func DOInit() *s3.Client {
	client := s3.New(s3.Options{
		Region: internal.REGION,
		Credentials: aws.NewCredentialsCache(
			credentials.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID:     config.DO_access_key_id,
					SecretAccessKey: config.DO_secret_access_key,
				},
			},
		),
		EndpointResolver: s3.EndpointResolverFromURL(config.DO_URL),
		UsePathStyle:     true,
	})

	if client == nil {
		log.Fatalln("Unable to initialize DigitalOcean S3 client")
	}
	log.Println("Initialized DigitalOcean S3 client")
	return client
}
