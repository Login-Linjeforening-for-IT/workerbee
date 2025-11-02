package services

import (
	"context"
	"mime/multipart"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsAlbums = map[string]string{
	"id":             "a.id",
	"name_no":        "a.name_no",
	"name_en":        "a.name_en",
	"description_no": "a.description_no",
	"description_en": "a.description_en",
	"year":           "a.year",
}

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

func (as *AlbumService) GetAlbums(ctx context.Context, orderBy, sort, limit_str, offset_str, search string) ([]models.AlbumsWithTotalCount, error) {
	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsAlbums)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return as.repo.GetAlbums(ctx, orderBySanitized, sortSanitized, search, limit, offset)
}
