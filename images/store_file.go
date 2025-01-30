package images

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type FileStore struct {
	baseDir string
}

var _ Store = &FileStore{}

func NewFileStore(baseDir string) *FileStore {
	return &FileStore{
		baseDir: baseDir,
	}
}

func (store *FileStore) GetImages(prefix string) ([]FileDetails, error) {
	path := filepath.Join(store.baseDir, prefix)

	dir, err := os.ReadDir(path)
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			return nil, &DirNotFoundError{Dir: prefix, Err: err}
		default:
			return nil, err
		}
	}

	files := make([]FileDetails, len(dir))
	for i, file := range dir {
		info, err := file.Info()
		if err != nil {
			return nil, err
		}

		files[i] = FileDetails{
			Name: file.Name(),
			Size: info.Size(),
			Path: filepath.Join(path, file.Name()),
		}
	}

	return files, nil
}

func (store *FileStore) UploadImage(dir string, id string, fileName string, file File) error {
	filePath := filepath.Join(store.baseDir, dir, id+fileName)

	out, err := os.Create(filePath)
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			return &DirNotFoundError{Dir: dir, Err: err}
		default:
			return err
		}
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
