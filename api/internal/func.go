package internal

import (
	"fmt"
	"net/url"
	"strings"
)

func DownscaleImage(width, height int) (int, int) {
	if width > MaxDimension || height > MaxDimension {
		if width >= height {
			ratio := float64(MaxDimension) / float64(width)
			return MaxDimension, int(float64(height) * ratio)
		} else {
			ratio := float64(MaxDimension) / float64(height)
			return int(float64(width) * ratio), MaxDimension
		}
	}
	return width, height
}

func ParseENAndNOArray(en, no []string) []map[string]string {
	var categories []map[string]string
	for i := range en {
		categories = append(categories, map[string]string{
			"en": en[i],
			"no": no[i],
		})
	}
	return categories
}

func ParseFromStringToSlice[T any](input string) ([]T, error) {
	if input != "" {
		slice, err := ParseCSVToSlice[T](input)
		if err != nil {
			return nil, ErrInvalid
		}
		return slice, nil
	} else {
		return make([]T, 0), nil
	}
}

func FormatNameWithCapitalFirstLetter(name string) string {
	if len(name) == 0 {
		return name
	}
	lowerName := strings.ToLower(name)
	return strings.ToUpper(string(lowerName[0])) + lowerName[1:]
}

func ParsePgArray(s string) []string {
	s = strings.Trim(s, "{}")
	if s == "" {
		return []string{}
	}
	return strings.Split(s, ",")
}

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

	parts := strings.SplitSeq(decoded, ",")
	for part := range parts {
		part = strings.TrimSpace(part)

		var v T
		if _, ok := any(v).(string); ok {
			v = any(part).(T)
		} else {
			_, err := fmt.Sscan(part, &v)
			if err != nil {
				return nil, ErrInvalid
			}
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

	return page * limit, limit, nil
}

func ParsePositiveInt(s string) (int, error) {
	var i int
	_, err := fmt.Sscan(s, &i)
	if err != nil || i < 0 {
		return 0, ErrInvalid
	}
	return i, nil
}
