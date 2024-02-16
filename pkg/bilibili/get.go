package bilibili

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiBase = "https://api.bilibili.com"

func getInfo(uri string, data interface{}) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v%v", apiBase, uri), nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return err
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var respData struct {
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
		TTL     int         `json:"ttl"`
	}
	if err := json.Unmarshal(raw, &respData); err != nil {
		return err
	}

	if respData.Code != 0 {
		return errors.New(respData.Message)
	}

	content, err := json.Marshal(respData.Data)
	if err != nil {
		return err
	}

	if data != nil {
		if err := json.Unmarshal(content, data); err != nil {
			return err
		}
	}

	return nil
}
