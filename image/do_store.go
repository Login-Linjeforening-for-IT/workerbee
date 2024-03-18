package image

type Store struct {
}

func NewDOStore() Store {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(server.config.DOKey, server.config.DOSecret, ""),
		Endpoint:         aws.String("https://ams3.digitaloceanspaces.com"),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String("ams3"),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
}
