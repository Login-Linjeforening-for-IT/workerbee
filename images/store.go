package images

import "io"

type Store interface {
	// Returns a list of images in the given directory.
	GetImages(dir string) ([]FileDetails, error)
	// Uploads the given file to the specified directory with the given ID and filename.
	UploadImage(dir string, fileName string, file File) error
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
