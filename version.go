package adminapi

import "os"

var version string = ""

func init() {
	if version == "" {
		v := os.Getenv("API_VERSION")
		if v != "" {
			version = v
			return
		}
		version = "unknown"
	}
}

func Version() string {
	return version
}
