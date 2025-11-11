package services

import (
	"context"
	"mime/multipart"
	"strconv"
	"time"
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
	repo  repositories.AlbumsRepository
	cache *CacheService
}

func NewAlbumService(repo repositories.AlbumsRepository, cache *CacheService) *AlbumService {
	return &AlbumService{
		repo:  repo,
		cache: cache,
	}
}

func (as *AlbumService) CreateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error) {
	album, err := as.repo.CreateAlbum(ctx, body)
	if err != nil {
		return models.CreateAlbum{}, err
	}

	pattern := "albums:*"
	as.cache.DeletePattern(ctx, pattern)

	return album, nil
}

func (as *AlbumService) UploadImagesToAlbum(ctx context.Context, id string, files []*multipart.FileHeader) error {
	err := as.repo.UploadImagesToAlbum(ctx, id, files)
	if err != nil {
		return err
	}

	as.cache.DeletePattern(ctx, "albums:*")

	return nil
}

func (as *AlbumService) GetAlbum(ctx context.Context, id string) (models.AlbumWithImages, error) {
	var album models.AlbumWithImages

	cacheKey := internal.AlbumKey(id)

	err := as.cache.GetJSON(ctx, cacheKey, &album)
	if err == nil {
		return album, nil
	}

	album, err = as.repo.GetAlbum(ctx, id)
	if err != nil {
		return models.AlbumWithImages{}, err
	}

	as.cache.Set(ctx, cacheKey, album, 5*time.Minute)
	return album, nil
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

	cacheKey := internal.AlbumsKey(orderBySanitized, sortSanitized, search, limit, offset)
	var albums []models.AlbumsWithTotalCount

	err = as.cache.GetJSON(ctx, cacheKey, &albums)
	if err == nil {
		return albums, nil
	}

	albums, err = as.repo.GetAlbums(ctx, orderBySanitized, sortSanitized, search, limit, offset)
	if err != nil {
		return nil, err
	}

	as.cache.Set(ctx, cacheKey, albums, 5*time.Minute)

	return albums, nil
}

func (as *AlbumService) UpdateAlbum(ctx context.Context, id string, body models.CreateAlbum) (models.CreateAlbum, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.CreateAlbum{}, internal.ErrInvalid
	}

	body.ID = idInt
	album, err := as.repo.UpdateAlbum(ctx, body)
	if err != nil {
		return models.CreateAlbum{}, err
	}

	as.cache.DeletePattern(ctx, "albums:*")

	return album, nil
}

func (as *AlbumService) DeleteAlbum(ctx context.Context, id string) (int, error) {
	respID, err := as.repo.DeleteAlbum(ctx, id)
	if err != nil {
		return 0, err
	}

	as.cache.DeletePattern(ctx, "albums:*")

	return respID, nil
}

func (as *AlbumService) DeleteAlbumImage(ctx context.Context, id string, imageName string) error {
	path := internal.ALBUM_PATH + id + "/" + imageName
	err := as.repo.DeleteAlbumImage(ctx, path, id)
	if err != nil {
		return err
	}

	as.cache.DeletePattern(ctx, "albums:*")

	return nil
}

func (as *AlbumService) SetAlbumCover(ctx context.Context, id string, imageName string) error {
	err := as.repo.SetAlbumCover(ctx, id, imageName)
	if err != nil {
		return err
	}

	as.cache.DeletePattern(ctx, "albums:*")

	return nil
}
