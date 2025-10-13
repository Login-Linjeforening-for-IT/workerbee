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
