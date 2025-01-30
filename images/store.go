package images

import "io"

type Store interface {
	GetImages(dir string) ([]FileDetails, error)
	UploadImage(dir string, id string, fileName string, file File) error
}

type FileDetails struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Path string `json:"filepath"`
}

type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}
