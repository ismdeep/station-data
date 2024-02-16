package blog

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// ShellHacks shell hacks
type ShellHacks struct {
}

// GetBloggerName get blogger name
func (receiver *ShellHacks) GetBloggerName() string {
	return "shellhacks"
}

// GetAllPageURLs get all page urls
func (receiver *ShellHacks) GetAllPageURLs() []string {
	pageURLs := []string{
		receiver.GetFirstPageURL(),
	}
	for i := 2; i <= 70; i++ {
		pageURLs = append(pageURLs, fmt.Sprintf("https://www.shellhacks.com/page/%v/", i))
	}

	return pageURLs
}

// GetFirstPageURL get first page url
func (receiver *ShellHacks) GetFirstPageURL() string {
	return "https://www.shellhacks.com/"
}

// GetLinksFromPage get links from page
func (receiver *ShellHacks) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//h2[@class="entry-title"]/a/@href`,
		`https://www.shellhacks.com/`)
}

// GetBlogInfo get blog info
func (receiver *ShellHacks) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="entry-title"]`,
		`//meta[@property="og:site_name"]/@content`,
		`//div[@class="entry-content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
