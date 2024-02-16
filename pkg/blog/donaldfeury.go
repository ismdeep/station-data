package blog

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// DonaldFeury struct
type DonaldFeury struct {
}

// GetBloggerName get blogger name
func (receiver *DonaldFeury) GetBloggerName() string {
	return "donaldfeury.xyz"
}

// GetAllPageURLs get all page urls
func (receiver *DonaldFeury) GetAllPageURLs() []string {
	links := make([]string, 0)
	links = append(links, "https://donaldfeury.xyz/")
	for i := 2; i <= 3; i++ {
		links = append(links, fmt.Sprintf("https://donaldfeury.xyz/page/%v/", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *DonaldFeury) GetFirstPageURL() string {
	return "https://donaldfeury.xyz/"
}

// GetLinksFromPage get blog links from page
func (receiver *DonaldFeury) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//div[@class="post-feed"]/article//a[@class="post-card-content-link"]/@href`,
		`https://donaldfeury.xyz`)
}

// GetBlogInfo get blog detail information
func (receiver *DonaldFeury) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="article-title"]`,
		`//h4[@class="author-name"]/a`,
		`//section[@class="gh-content gh-canvas"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
