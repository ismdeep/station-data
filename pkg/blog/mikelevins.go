package blog

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// MikelEvins struct
type MikelEvins struct {
}

// GetBloggerName get blogger name
func (receiver *MikelEvins) GetBloggerName() string {
	return "mikelevins.github.io"
}

// GetAllPageURLs get all page urls
func (receiver *MikelEvins) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	for i := 2; i <= 26; i++ {
		links = append(links, fmt.Sprintf("https://mikelevins.github.io/posts/page/%v/", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *MikelEvins) GetFirstPageURL() string {
	return "https://mikelevins.github.io/posts/"
}

// GetLinksFromPage get blog links from page
func (receiver *MikelEvins) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//div[@class="post-heading"]/h1/a/@href`,
		``)
}

// GetBlogInfo get blog detail information
func (receiver *MikelEvins) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	v, err := GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//div[@class="post-heading"]/h1`,
		`//div[@class="sidebar-content"]/p`,
		`//div[@class="post"]`,
		func(doc *html.Node) (time.Time, error) {
			sNode := htmlquery.FindOne(doc, `//span[@class="post-date"]/a`)
			if sNode == nil {
				return time.Now(), errors.New("time not found")
			}
			s := htmlquery.InnerText(sNode)
			s = strings.ReplaceAll(s, "# ", "")
			return time.Parse("Jan 2, 2006", s)
		})
	if err != nil {
		return nil, err
	}

	v.Author = strings.TrimSpace(strings.ReplaceAll(v.Author, "by ", ""))
	return v, nil
}
