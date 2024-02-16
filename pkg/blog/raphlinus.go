package blog

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// RaphLinus struct
type RaphLinus struct {
}

// GetBloggerName get blogger name
func (receiver *RaphLinus) GetBloggerName() string {
	return "raphlinus.github.io"
}

// GetAllPageURLs get all page urls
func (receiver *RaphLinus) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	return links
}

// GetFirstPageURL get first page url
func (receiver *RaphLinus) GetFirstPageURL() string {
	return "https://raphlinus.github.io/"
}

// GetLinksFromPage get blog links from page
func (receiver *RaphLinus) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//ul[@class="post-list"]//a[@class="post-link"]/@href`,
		`https://raphlinus.github.io`)
}

// GetBlogInfo get blog detail information
func (receiver *RaphLinus) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="post-title p-name"]`,
		`//span[@class="username"]`,
		`//div[@class="post-content e-content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
