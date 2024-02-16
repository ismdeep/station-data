//go:build !test

package youtube

import (
	"io"
	"net/http"
	"time"
)

// GetYouTubeHTML get youtube html content
func GetYouTubeHTML(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := (&http.Client{
		Timeout: 10 * time.Second,
	}).Do(req)
	if err != nil {
		return "", err
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	content := string(raw)
	return content, nil
}
