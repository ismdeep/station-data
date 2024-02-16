package blog

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// Cybartist struct
type Cybartist struct {
}

// GetBloggerName get blogger name
func (receiver *Cybartist) GetBloggerName() string {
	return "cybart.ist"
}

// GetAllPageURLs get all page urls
func (receiver *Cybartist) GetAllPageURLs() []string {
	return []string{"https://cybart.ist/archive.html"}
}

// GetFirstPageURL get first page url
func (receiver *Cybartist) GetFirstPageURL() string {
	return "https://cybart.ist/archive.html"
}

// GetLinksFromPage get blog links from page
func (receiver *Cybartist) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//li[@class="item"]//a[@class="item__header"]/@href`, "https://cybart.ist")
}

// GetBlogInfo get blog detail information
func (receiver *Cybartist) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//div[@class="article__header"]/header/h1`,
		`//meta[@itemprop="author"]/@content`,
		`//div[@class="article__content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//div[@class="article__info clearfix"]/ul[@class="right-col menu"]/li/span`, "Jan 02, 2006")
		})
}
