package services

import (
	"context"
	"strconv"
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

func (as *AlbumService) UploadImagesToAlbum(ctx context.Context, id string, uploads []models.UploadImages) ([]models.UploadPictureResponse, error) {
	return as.repo.UploadImagesToAlbum(ctx, id, uploads)
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

func (as *AlbumService) UpdateAlbum(ctx context.Context, id string, body models.CreateAlbum) (models.CreateAlbum, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.CreateAlbum{}, internal.ErrInvalid
	}

	body.ID = idInt
	return as.repo.UpdateAlbum(ctx, body)
}

func (as *AlbumService) DeleteAlbum(ctx context.Context, id string) (int, error) {
	return as.repo.DeleteAlbum(ctx, id)
}

func (as *AlbumService) DeleteAlbumImage(ctx context.Context, id string, imageName string) error {
	path := internal.ALBUM_PATH + id + "/" + imageName
	return as.repo.DeleteAlbumImage(ctx, path, id)
}

func (as *AlbumService) SetAlbumCover(ctx context.Context, id string, imageName string) error {
	return as.repo.SetAlbumCover(ctx, id, imageName)
}
