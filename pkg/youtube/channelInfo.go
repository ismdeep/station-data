package youtube

import (
	"bytes"
	"errors"
	"strings"

	"github.com/antchfx/htmlquery"
)

// ChannelInfo global info
type ChannelInfo struct {
	ChannelID string
	Name      string
}

// GetChannelInfo get youtube global info by global home url
func GetChannelInfo(channelHomeURL string) (*ChannelInfo, error) {
	content, err := GetYouTubeHTML(channelHomeURL)
	if err != nil {
		return nil, err
	}

	doc, err := htmlquery.Parse(bytes.NewReader([]byte(content)))
	if err != nil {
		return nil, err
	}

	channelIDNode := htmlquery.FindOne(doc, `//meta[@itemprop="identifier"]/@content`)
	if channelIDNode == nil {
		return nil, errors.New("global id not found")
	}

	nameNode := htmlquery.FindOne(doc, `//meta[@property="og:title"]/@content`)
	if nameNode == nil {
		return nil, errors.New("name is not found")
	}

	return &ChannelInfo{
		ChannelID: strings.TrimSpace(htmlquery.InnerText(channelIDNode)),
		Name:      strings.TrimSpace(htmlquery.InnerText(nameNode)),
	}, nil
}
