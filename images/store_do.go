package images

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type DOStore struct {
	client *s3.S3
}

var _ Store = &DOStore{}

func NewDOStore(
	doKey string,
	doSecret string,
	// TODO: More config
) (*DOStore, error) {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(doKey, doSecret, ""),
		Endpoint:         aws.String("https://ams3.digitaloceanspaces.com"),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String("ams3"),
	}

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(sess)
	return &DOStore{
		client: s3Client,
	}, nil
}

func (store *DOStore) GetImages(dir string) ([]FileDetails, error) {
	return []FileDetails{}, nil
}

func (store *DOStore) UploadImage(dir string, name string, file File) error {
	return nil
}
