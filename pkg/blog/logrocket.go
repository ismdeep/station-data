package blog

import (
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// LogRocket struct
type LogRocket struct {
}

// GetBloggerName get blogger name
func (receiver *LogRocket) GetBloggerName() string {
	return "logrocket"
}

// GetAllPageURLs get all page urls
func (receiver *LogRocket) GetAllPageURLs() []string {
	links := []string{"https://blog.logrocket.com/"}
	for i := 2; i <= 240; i++ {
		links = append(links, fmt.Sprintf("https://blog.logrocket.com/page/%v/", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *LogRocket) GetFirstPageURL() string {
	return "https://blog.logrocket.com/"
}

// GetLinksFromPage get blog links from page
func (receiver *LogRocket) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//div[@class="card-block"]//h2[@class="card-title"]/a/@href`, "")
}

// GetBlogInfo get blog detail information
func (receiver *LogRocket) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="posttitle"]`,
		`//div[@class="mainheading"]//a[@class="text-capitalize link-dark"]/text()`,
		`//article[@class="article-post"]`,
		func(doc *html.Node) (time.Time, error) {
			s := htmlquery.OutputHTML(doc, true)
			s = s[strings.Index(s, `<meta property="article:published_time" content="`)+len(`<meta property="article:published_time" content="`):]
			s = s[:strings.Index(s, `"`)]
			return time.Parse(time.RFC3339, s)
		})
}
