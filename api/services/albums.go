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

func (as *AlbumService) CreateAlbum(ctx context.Context, images []*multipart.FileHeader, body models.Album) error {
	return as.repo.CreateAlbum(ctx, images, body)
}
