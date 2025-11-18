package handlers

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var purgeKeySanitizer = regexp.MustCompile(`[^a-zA-Z0-9_\-:.]`)

func sanitizeKey(key string) string {
	k := strings.TrimSpace(key)
	if k == "" {
		return ""
	}

	k = strings.ReplaceAll(k, "/", "_")
	k = strings.ReplaceAll(k, " ", "-")
	k = purgeKeySanitizer.ReplaceAllString(k, "")
	return k
}

func SetSurrogatePurgeHeader(c *gin.Context, keys ...string) {
	set := map[string]struct{}{}
	for _, key := range keys {
		if s := sanitizeKey(key); s != "" {
			set[s] = struct{}{}
		}
	}

	if len(set) == 0 {
		return
	}

	out := make([]string, 0, len(set))
	for k := range set {
		out = append(out, k)
	}

	c.Header("Surrogate-Key", strings.Join(out, "|"))
}
