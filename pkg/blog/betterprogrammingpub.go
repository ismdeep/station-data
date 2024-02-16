package blog

import (
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// BetterProgrammingPub struct
type BetterProgrammingPub struct {
}

// GetBloggerName get blogger name
func (receiver *BetterProgrammingPub) GetBloggerName() string {
	return "betterprogramming.pub"
}

// GetAllPageURLs get all page urls
func (receiver *BetterProgrammingPub) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	return links
}

// GetFirstPageURL get first page url
func (receiver *BetterProgrammingPub) GetFirstPageURL() string {
	return "https://betterprogramming.pub/"
}

// GetLinksFromPage get blog links from page
func (receiver *BetterProgrammingPub) GetLinksFromPage(pageURL string) ([]string, error) {
	links, err := GetBlogLinks(pageURL,
		`//article//a[@aria-label="Post Preview Title"]/@href`,
		`https://betterprogramming.pub`)
	if err != nil {
		return nil, err
	}

	var result []string

	for i := 0; i < len(links); i++ {
		s := links[i]
		if strings.Index(s, "?source=") > 0 {
			result = append(result, s[:strings.Index(s, "?source=")])
		}
	}
	return result, nil
}

// GetBlogInfo get blog detail information
func (receiver *BetterProgrammingPub) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@name="author"]/@content`,
		`//section`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
