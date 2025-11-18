package middleware

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var purgeKeySanitizer = regexp.MustCompile(`[^a-zA-Z0-9_\-:.]`)

var affectedKeys = map[string]struct {
	modifying    []string
	nonModifying []string
}{
	"jobs": {
		modifying:    []string{"jobs"},
		nonModifying: []string{"jobs"},
	},
	"organizations": {
		modifying:    []string{"organizations"},
		nonModifying: []string{"organizations", "jobs"},
	},
	// add more paths here
}

// SetSurrogatePurgeHeader sets the Surrogate-Purge header for cache purging.
func SetHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		trimmed := strings.Split(c.Request.URL.Path, "/")

		prefix := trimmed[2]

		affectedKeys := getAffectedKeys(prefix, c.Request.Method)

		setSurrogateHeader(c, affectedKeys...)

		c.Next()
	}
}

func getAffectedKeys(path, method string) []string {
	keys, ok := affectedKeys[path]
	if !ok {
		return []string{}
	}

	switch method {
	case http.MethodGet, http.MethodHead:
		return keys.nonModifying
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		return keys.modifying
	default:
		return []string{}
	}
}

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

func setSurrogateHeader(c *gin.Context, keys ...string) {
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
