package blog

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// Johnj struct
type Johnj struct {
}

// GetBloggerName get blogger name
func (receiver *Johnj) GetBloggerName() string {
	return "johnj.com"
}

// GetAllPageURLs get all page urls
func (receiver *Johnj) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	return links
}

// GetFirstPageURL get first page url
func (receiver *Johnj) GetFirstPageURL() string {
	return "http://johnj.com/posts/"
}

// GetLinksFromPage get blog links from page
func (receiver *Johnj) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//main[@class="content"]/article[@class="post-snippet"]/a[@class="post-title"]/@href`,
		`http://johnj.com`)
}

// GetBlogInfo get blog detail information
func (receiver *Johnj) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//main[@class="content"]/h1`,
		`//header[@class="site-header"]/nav[@class="site-nav"]/a[@class="logo"]`,
		`//main[@class="content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//time[@class="postdate"]/@datetime`, time.RFC3339)
		})
}
