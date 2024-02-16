package blog

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// Qiwsir struct
type Qiwsir struct {
}

// GetBloggerName get blogger name
func (receiver *Qiwsir) GetBloggerName() string {
	return "qiwsir.github.io"
}

// GetAllPageURLs get all page urls
func (receiver *Qiwsir) GetAllPageURLs() []string {
	var links []string
	links = append(links, receiver.GetFirstPageURL())
	for i := 2; i <= 16; i++ {
		links = append(links, fmt.Sprintf("https://qiwsir.github.io/page/%v/", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *Qiwsir) GetFirstPageURL() string {
	return "https://qiwsir.github.io/"
}

// GetLinksFromPage get links from page
func (receiver *Qiwsir) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//main[@class="app-body"]/article[@class="article-card"]/h2[@class="article-head"]/a/@href`,
		`https://qiwsir.github.io`)
}

// GetBlogInfo get blog info
func (receiver *Qiwsir) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//article[@class="post-article"]/h2`,
		`//div[@class="header-container"]/a[@class="home-link"]/span`,
		`//article[@class="post-article"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//p[@class="post-date"]`, "2006-01-02")
		})
}
