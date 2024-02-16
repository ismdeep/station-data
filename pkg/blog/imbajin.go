package blog

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// Imbajin struct
type Imbajin struct {
}

// GetBloggerName get blogger name
func (receiver *Imbajin) GetBloggerName() string {
	return "imbajin"
}

// GetAllPageURLs get all page urls
func (receiver *Imbajin) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	for i := 2; i <= 13; i++ {
		links = append(links, fmt.Sprintf("http://www.imbajin.com/archives/page/%v/", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *Imbajin) GetFirstPageURL() string {
	return "http://www.imbajin.com/archives/"
}

// GetLinksFromPage get blog links from page
func (receiver *Imbajin) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//li[@class="post-item"]/span/a/@href`,
		`http://www.imbajin.com`)
}

// GetBlogInfo get blog detail information
func (receiver *Imbajin) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//span[@class="author"]/span[@itemprop="name"]`,
		`//div[@class="content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="og:updated_time"]/@content`, time.RFC3339)
		})
}
