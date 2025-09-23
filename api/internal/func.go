package internal

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
