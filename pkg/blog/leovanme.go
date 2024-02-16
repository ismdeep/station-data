package blog

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// LeovanMe struct
type LeovanMe struct {
}

// GetBloggerName get blogger name
func (receiver *LeovanMe) GetBloggerName() string {
	return "leovan.me"
}

// GetAllPageURLs get all page urls
func (receiver *LeovanMe) GetAllPageURLs() []string {
	return []string{"https://leovan.me/cn/"}
}

// GetFirstPageURL get first page url
func (receiver *LeovanMe) GetFirstPageURL() string {
	return "https://leovan.me/cn/"
}

// GetLinksFromPage get blog links from page
func (receiver *LeovanMe) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//p[@class="list-page"]/a/@href`, "https://leovan.me")
}

// GetBlogInfo get blog detail information
func (receiver *LeovanMe) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="title"]`,
		`//meta[@name="author"]/@content`,
		`//article[@class="main"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@name="date"]/@content`, "2006-01-02")
		})
}
