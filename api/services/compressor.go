package services

import (
	"bytes"
	"context"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
	"workerbee/internal"
	"workerbee/repositories"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

type CompressorService struct {
	imageRepo  repositories.ImageRepository
	albumsRepo repositories.AlbumsRepository
}

func initCompressor(cmpr *CompressorService) {
	ctx := context.Background()
	go func() {
		ticker := time.NewTicker(1 * time.Hour)

		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				log.Println("Compressing albums...")
				err := cmpr.CompressAllAlbums(ctx)
				if err != nil {
					log.Printf("Error compressing albums: %v\n", err)
				} else {
					log.Println("Successfully compressed album images")
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func NewCompressorService(imageRepo repositories.ImageRepository, albumsRepo repositories.AlbumsRepository) *CompressorService {
	compressor := &CompressorService{
		imageRepo:  imageRepo,
		albumsRepo: albumsRepo,
	}

	initCompressor(compressor)

	return compressor
}

func isRawKey(filename string) bool {
	split := strings.Split(filename, "_")
	if len(split) < 2 {
		return false
	}
	return split[2] == "raw"
}

func compressedKeyFromRaw(rawKey string) string {
	newKey := strings.Replace(rawKey, "_raw_", "_", 1)

	ext := filepath.Ext(newKey)
	if ext != "" {
		newKey = strings.TrimSuffix(newKey, ext) + ".webp"
	} else if !strings.HasSuffix(newKey, ".webp") {
		newKey += ".webp"
	}

	if !strings.HasPrefix(newKey, internal.ALBUM_PATH) {
		newKey = internal.ALBUM_PATH + newKey
	}

	return newKey
}

func (cs *CompressorService) CompressAllAlbums(ctx context.Context) error {
	images, err := cs.imageRepo.GetImagesInPath(ctx, internal.ALBUM_PATH)
	if err != nil {
		return err
	}

	rawImages := []string{}
	for _, img := range images {
		if isRawKey(img) {
			rawImages = append(rawImages, img)
		}
	}

	for _, rel := range rawImages {
		albumParts := strings.Split(rel, "/")
		if len(albumParts) < 2 {
			continue
		}
		albumID := albumParts[0]
		rawFilename := albumParts[1]
		rawKey := internal.ALBUM_PATH + rel

		rc, _, err := cs.imageRepo.GetObject(ctx, rawKey)
		if err != nil {
			return err
		}

		buf, err := cs.compressImageStream(rc)
		rc.Close()
		if err != nil {
			return err
		}

		compressedKey := compressedKeyFromRaw(rel)
		if err := cs.imageRepo.UploadImage(ctx, compressedKey, "image/webp", buf); err != nil {
			return err
		}

		err = cs.albumsRepo.DeleteAlbumImage(ctx, internal.ALBUM_PATH+albumID+"/"+rawFilename, albumID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cs *CompressorService) compressImageStream(src io.Reader) (*bytes.Buffer, error) {
	img, err := imaging.Decode(src)
	if err != nil {
		return nil, err
	}

	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	newW, newH := internal.DownscaleImage(w, h)
	if newW != w || newH != h {
		img = imaging.Resize(img, newW, newH, imaging.Lanczos)
	}

	buf := new(bytes.Buffer)
	if err := webp.Encode(buf, img, &webp.Options{Lossless: false, Quality: 80}); err != nil {
		return nil, err
	}

	return buf, nil
}
