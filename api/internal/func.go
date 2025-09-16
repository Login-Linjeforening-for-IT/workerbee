package internal

func SanitizeSort(col, dir string, allowed map[string]string) (string, string, error) {
	c, ok := allowed[col]
	if !ok {
		return "", "", ErrInvalidSort
	}
	if dir != "asc" && dir != "desc" {
		return "", "", ErrInvalidSort
	}
	return c, dir, nil
}
