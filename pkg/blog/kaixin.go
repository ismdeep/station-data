package blog

import (
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// Kaixin struct
type Kaixin struct {
}

// GetBloggerName get blogger name
func (receiver *Kaixin) GetBloggerName() string {
	return "kaix.in"
}

// GetAllPageURLs get all page urls
func (receiver *Kaixin) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	for i := 2; i <= 8; i++ {
		links = append(links, fmt.Sprintf("https://kaix.in/0001/more/%v/urls.json", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *Kaixin) GetFirstPageURL() string {
	return "https://kaix.in/more/urls.json"
}

// GetLinksFromPage get blog links from page
func (receiver *Kaixin) GetLinksFromPage(pageURL string) ([]string, error) {
	content, err := GetHTML(pageURL)
	if err != nil {
		return nil, err
	}

	type RespData struct {
		Pages []struct {
			URL string `json:"url"`
		} `json:"pages"`
	}

	var respData RespData
	if err := json.Unmarshal([]byte(content), &respData); err != nil {
		return nil, err
	}

	var links []string
	for _, v := range respData.Pages {
		links = append(links, "https://kaix.in"+v.URL)
	}

	return links, nil
}

// GetBlogInfo get blog detail information
func (receiver *Kaixin) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@name="author"]/@content`,
		`//div[@itemprop="articleBody"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
