package blog

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// MorningSpace struct
type MorningSpace struct {
}

// GetBloggerName get blogger name
func (receiver *MorningSpace) GetBloggerName() string {
	return "morningspace"
}

// GetAllPageURLs get all page urls
func (receiver *MorningSpace) GetAllPageURLs() []string {
	items := make([]string, 0)
	items = append(items, "https://morningspace.github.io/")
	for i := 2; i <= 14; i++ {
		items = append(items, fmt.Sprintf("https://morningspace.github.io/page%v/", i))
	}
	return items
}

// GetFirstPageURL get first page url
func (receiver *MorningSpace) GetFirstPageURL() string {
	return "https://morningspace.github.io/"
}

// GetLinksFromPage get blog links from page
func (receiver *MorningSpace) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//div[@class="entries-list"]/div[@class="list__item"]/article[@class="archive__item"]/h2[@class="archive__item-title no_toc"]/a/@href`,
		`https://morningspace.github.io`)
}

// GetBlogInfo get blog detail information
func (receiver *MorningSpace) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="page__title"]`,
		`//div[@class="author__content"]/h3[@class="author__name p-name"]/a`,
		`//article[@class="page h-entry"]/div[@class="page__inner-wrap"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//article[@class="page h-entry"]/meta[@itemprop="datePublished"]/@content`, time.RFC3339)
		})
}
