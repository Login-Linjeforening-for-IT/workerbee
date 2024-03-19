package image

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Store interface {
	Requester
}

type DOStore struct {
	DO_client *s3.S3
	*Requester
}

func NewDOStore(DOConfig *aws.Config) *DOStore {
	new_session, err := session.NewSession(DOConfig)
	if err != nil {
		fmt.Println("Could not create new session!\n")
	}

	return &DOStore{
		DO_client: s3.New(new_session),
	}
}
