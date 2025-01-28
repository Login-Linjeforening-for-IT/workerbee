package images

type DirNotFoundError struct {
	Dir string

	Err error
}

func (e *DirNotFoundError) Error() string {
	return "directory not found: " + e.Dir + ": " + e.Err.Error()
}

func (e *DirNotFoundError) Unwrap() error {
	return e.Err
}
