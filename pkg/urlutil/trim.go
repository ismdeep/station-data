package urlutil

import "strings"

func TrimPath(uri string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(uri, "https://", "{https-schema}"),
					"http://", "{http-schema}"),
				"//", "/"),
			"{http-schema}", "http://"),
		"{https-schema}", "https://")
}
