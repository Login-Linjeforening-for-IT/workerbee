package sessionstore

import "fmt"

type NotFoundError struct {
	Key      string
	Resource string

	Err error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found in %s", e.Key, e.Resource)
}

func (e *NotFoundError) Unwrap() error {
	return e.Err
}
