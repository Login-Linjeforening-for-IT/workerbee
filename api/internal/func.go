package internal

import (
	"fmt"
	"net/url"
	"strings"
)

func SanitizeSort(col, dir string, allowed map[string]string) (string, string, error) {
	c, ok := allowed[col]
	if !ok {
		return "", "", ErrInvalid
	}
	if dir != "asc" && dir != "desc" {
		return "", "", ErrInvalid
	}
	return c, dir, nil
}

func ParseCSVToSlice[T any](s string) ([]T, error) {
	var result []T
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return nil, ErrInvalid
	}
	parts := strings.Split(decoded, ",")
	for _, part := range parts {
		var v T
		_, err := fmt.Sscan(part, &v)
		if err != nil {
			return nil, ErrInvalid
		}
		result = append(result, v)
	}

	return result, nil
}

func CalculateOffset(page_str, limit_str string) (int, int, error) {
	page, err := ParsePositiveInt(page_str)
	if err != nil {
		return 0, 0, ErrInvalid
	}

	limit, err := ParsePositiveInt(limit_str)
	if err != nil {
		return 0, 0, ErrInvalid
	}

	if limit <= 0 || limit > 50 {
		limit = 20
	}

	if page < 0 {
		page = 0
	}

	return page*limit, limit, nil
}

func ParsePositiveInt(s string) (int, error) {
	var i int
	_, err := fmt.Sscan(s, &i)
	if err != nil || i < 0 {
		return 0, ErrInvalid
	}
	return i, nil
}
