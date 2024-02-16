package blog

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// Izsk struct
type Izsk struct {
}

// GetBloggerName get blogger name
func (receiver *Izsk) GetBloggerName() string {
	return "izsk.me"
}

// GetAllPageURLs get all page urls
func (receiver *Izsk) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	for i := 2; i <= 19; i++ {
		links = append(links, fmt.Sprintf("https://izsk.me/page/%v/", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *Izsk) GetFirstPageURL() string {
	return "https://izsk.me/"
}

// GetLinksFromPage get blog links from page
func (receiver *Izsk) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//article[@class="post post-type-normal "]//a[@class="post-title-link"]/@href`,
		`https://izsk.me/`)
}

// GetBlogInfo get blog detail information
func (receiver *Izsk) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@property="article:author"]/@content`,
		`//div[@class="post-body"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//time[@itemprop="dateCreated datePublished"]/@datetime`, time.RFC3339)
		})
}
