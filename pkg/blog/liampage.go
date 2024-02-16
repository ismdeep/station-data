package blog

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// LiamPage struct
type LiamPage struct {
}

// GetBloggerName get blogger name
func (receiver *LiamPage) GetBloggerName() string {
	return "liam.page"
}

// GetAllPageURLs get all page urls
func (receiver *LiamPage) GetAllPageURLs() []string {
	links := make([]string, 0)
	links = append(links, receiver.GetFirstPageURL())
	for i := 2; i <= 76; i++ {
		links = append(links, fmt.Sprintf("https://liam.page/page/%v/", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *LiamPage) GetFirstPageURL() string {
	return "https://liam.page/"
}

// GetLinksFromPage get blog links from page
func (receiver *LiamPage) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//div[@class="content index posts-expand"]/article[@class="post-block"]//h2[@class="post-title"]/a/@href`,
		`https://liam.page`)
}

// GetBlogInfo get blog detail information
func (receiver *LiamPage) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="post-title"]`,
		`//span[@itemprop="author"]/meta[@itemprop="name"]/@content`,
		`//div[@class="post-body"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
