package middleware

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

var purgeKeySanitizer = regexp.MustCompile(`[^a-zA-Z0-9_\-:.]`)

var affectedKeys = map[string]struct {
	modifying    []string
	nonModifying []string
}{
	"jobs": {
		modifying:    []string{"jobs", "stats"},
		nonModifying: []string{"jobs"},
	},
	"organizations": {
		modifying:    []string{"organizations", "jobs", "stats", "events"},
		nonModifying: []string{"organizations"},
	},
	"rules": {
		modifying:    []string{"rules", "events", "stats"},
		nonModifying: []string{"rules"},
	},
	"audiences": {
		modifying:    []string{"audiences", "jobs", "events", "stats"},
		nonModifying: []string{"audiences"},
	},
	"locations": {
		modifying:    []string{"locations", "jobs", "events", "stats"},
		nonModifying: []string{"locations"},
	},
	"albums": {
		modifying:    []string{"albums"},
		nonModifying: []string{"albums"},
	},
	"alerts": {
		modifying:    []string{"alerts"},
		nonModifying: []string{"alerts"},
	},
	"stats": {
		modifying:    []string{"stats"},
		nonModifying: []string{"stats"},
	},
	"events": {
		modifying:    []string{"events", "stats", "categories", "audiences", "locations", "rules", "organizations"},
		nonModifying: []string{"events"},
	},
	"categories": {
		modifying:    []string{"categories", "events", "stats", "jobs"},
		nonModifying: []string{"categories"},
	},
	"honeys": {
		modifying:    []string{"honeys", "stats", "text"},
		nonModifying: []string{"honeys"},
	},
	"text": {
		modifying:    []string{"text", "stats", "honeys"},
		nonModifying: []string{"text"},
	},
}

// SetSurrogatePurgeHeader sets the Surrogate-Purge header for cache purging.
func SetHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		routePattern := c.FullPath()
		if routePattern == "" {
			c.Next()
			return
		}

		parts := strings.Split(strings.Trim(routePattern, "/"), "/")
		resource := ""
		for _, p := range parts {
			if p != "" && !strings.Contains(internal.BASE_PATH, p) {
				resource = p
				break
			}
		}

		affectedKeys := getAffectedKeys(resource, c.Request.Method)

		var affectedKeysList = make([]string, 0, len(c.Params))
		for _, param := range c.Params {
			affectedKeysList = append(affectedKeysList, fmt.Sprintf("%s:%s", resource, param.Value))
		}

		if len(affectedKeysList) > 0 {
			affectedKeys = append(affectedKeys, affectedKeysList...)
		}

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
