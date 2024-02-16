package blog

import (
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"

	"github.com/ismdeep/station-data/internal/schema"
)

// Solidot struct
type Solidot struct {
}

// GetBloggerName get blogger name
func (receiver *Solidot) GetBloggerName() string {
	return "solidot"
}

// GetAllPageURLs get all page urls
func (receiver *Solidot) GetAllPageURLs() []string {
	links := []string{"https://www.solidot.org"}
	t := time.Now()
	stop, _ := time.Parse("20060102", "20150101")
	for {
		t = t.AddDate(0, 0, -1)
		links = append(links, fmt.Sprintf("https://www.solidot.org/?issue=%v", t.Format("20060102")))
		if t.Unix() < stop.Unix() {
			break
		}
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *Solidot) GetFirstPageURL() string {
	return "https://www.solidot.org"
}

// GetLinksFromPage get blog links from page
func (receiver *Solidot) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//div[@class="block_m"]//h2/a/@href`, "https://www.solidot.org")
}

// GetBlogInfo get blog detail information
func (receiver *Solidot) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//div[@class="ct_tittle"]/div[@class="bg_htit"]/h2`,
		`//div[@class="talk_time"]/a`,
		`//div[@class="p_mainnew"]`,
		func(doc *html.Node) (time.Time, error) {
			s := htmlquery.OutputHTML(doc, true)
			s = s[strings.Index(s, `"pubDate": "`)+len(`"pubDate": "`):]
			s = s[:strings.Index(s, `"`)]
			return time.Parse("2006-01-02T15:04:05", s)
		})
}
