package blog

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// IOMeter model
type IOMeter struct {
}

// GetBloggerName get blog name
func (receiver *IOMeter) GetBloggerName() string {
	return "io-meter"
}

// GetAllPageURLs get all page urls
func (receiver *IOMeter) GetAllPageURLs() []string {
	var list []string
	list = append(list, receiver.GetFirstPageURL())
	for i := 2; i <= 5; i++ {
		list = append(list, fmt.Sprintf("https://io-meter.com/archives/page/%v/", i))
	}
	return list
}

// GetFirstPageURL get first page url
func (receiver *IOMeter) GetFirstPageURL() string {
	return "https://io-meter.com/archives/"
}

// GetLinksFromPage get links from page
func (receiver *IOMeter) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//article/div[@class="post-content"]/header/h1[@class="title"]/a/@href`, "https://io-meter.com")
}

// GetBlogInfo get blog info
func (receiver *IOMeter) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@name="author"]/@content`,
		`//div[@class="post-content"]/div[@class="entry"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//time/@datetime`, time.RFC3339)
		})
}
