package blog

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// KeldosMe blogger
type KeldosMe struct {
}

// GetBloggerName get blogger name
func (receiver *KeldosMe) GetBloggerName() string {
	return "blog.keldos.me"
}

// GetAllPageURLs get all page url
func (receiver *KeldosMe) GetAllPageURLs() []string {
	return []string{
		receiver.GetFirstPageURL(),
		"https://blog.keldos.me/page/2/",
	}
}

// GetFirstPageURL get first page url
func (receiver *KeldosMe) GetFirstPageURL() string {
	return "https://blog.keldos.me/"
}

// GetLinksFromPage get link from page
func (receiver *KeldosMe) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//article//h2[@class="post-title"]/a/@href`, `https://blog.keldos.me`)
}

// GetBlogInfo get blog info
func (receiver *KeldosMe) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@property="article:author"]/@content`,
		`//div[@class="content post posts-expand"]//div[@class="post-body"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
