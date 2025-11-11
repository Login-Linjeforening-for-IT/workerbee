package internal

import "fmt"

func AlbumKey(id any) string {
	return fmt.Sprintf("albums:%v", id)
}

func AlbumsKey(orderBy, sort, search string, limit, offset int) string {
	return fmt.Sprintf("albums:%s:%s:%s:%d:%d", orderBy, sort, search, limit, offset)
}
