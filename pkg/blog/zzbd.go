package blog

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// Zzbd struct
type Zzbd struct {
}

// GetBloggerName get blogger name
func (receiver *Zzbd) GetBloggerName() string {
	return "blog.zzbd.org"
}

// GetAllPageURLs get all page urls
func (receiver *Zzbd) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	for i := 2; i <= 27; i++ {
		links = append(links, fmt.Sprintf("https://blog.zzbd.org/page/%v/", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *Zzbd) GetFirstPageURL() string {
	return "https://blog.zzbd.org/"
}

// GetLinksFromPage get blog links from page
func (receiver *Zzbd) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//article[@class="post post-type-normal"]//h1[@class="post-title"]/a[@class="post-title-link"]/@href`,
		`https://blog.zzbd.org`)
}

// GetBlogInfo get blog detail information
func (receiver *Zzbd) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="post-title"]`,
		`//meta[@property="article:author"]/@content`,
		`//div[@class="post-body"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
