package services

import (
	"context"
	"mime/multipart"
	"workerbee/models"
	"workerbee/repositories"
)

type AlbumService struct {
	repo repositories.AlbumsRepository
}

func NewAlbumService(repo repositories.AlbumsRepository) *AlbumService {
	return &AlbumService{
		repo: repo,
	}
}

func (as *AlbumService) CreateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error) {
	return as.repo.CreateAlbum(ctx, body)
}

func (as *AlbumService) UploadImagesToAlbum(ctx context.Context, id string, files []*multipart.FileHeader) error {
	return as.repo.UploadImagesToAlbum(ctx, id, files)
}

func (as *AlbumService) GetAlbum(ctx context.Context, id string) (models.AlbumWithImages, error) {
	return as.repo.GetAlbum(ctx, id)
}
