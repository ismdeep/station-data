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
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

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
