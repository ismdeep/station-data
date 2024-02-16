package youtube

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// GetAllVideos get global all videos
func GetAllVideos(channelID string) ([]string, error) {
	url := fmt.Sprintf("https://www.youtube.com/channel/%v", channelID)
	content, err := GetYouTubeHTML(url)
	if err != nil {
		return nil, err
	}

	videoIDMap := make(map[string]bool)
	const flag = `"videoId":"`
	for {
		if i := strings.Index(content, flag); i >= 0 {
			content = content[i+len(flag):]
			videoID := content[:strings.Index(content, `"`)]
			videoIDMap[videoID] = true
			continue
		}
		break
	}

	var videoIDList []string
	for videoID := range videoIDMap {
		videoIDList = append(videoIDList, videoID)
	}

	return videoIDList, nil
}

// GetAllVideosFromPlaylist get all videos from playlist
func GetAllVideosFromPlaylist(playlistID string) (string, []string, error) {
	url := fmt.Sprintf("https://www.youtube.com/playlist?list=%v", playlistID)
	content, err := GetYouTubeHTML(url)
	if err != nil {
		return "", nil, err
	}

	ct := content
	ci := strings.Index(ct, `,"channelId":"`)
	if ci < 0 {
		return "", nil, errors.New("not found channelId")
	}
	channelID := ct[ci+14:]
	channelID = channelID[:strings.Index(channelID, `"`)]

	videoIDMap := make(map[string]bool)
	const flag = `"videoId":"`
	for {
		if i := strings.Index(content, flag); i >= 0 {
			content = content[i+len(flag):]
			videoID := content[:strings.Index(content, `"`)]
			videoIDMap[videoID] = true
			continue
		}
		break
	}

	var videoIDList []string
	for videoID := range videoIDMap {
		videoIDList = append(videoIDList, videoID)
	}

	return channelID, videoIDList, nil
}

// VideoInfo video info
type VideoInfo struct {
	VideoID       string
	Link          string
	ChannelID     string
	Name          string
	DatePublished string
	Author        string
}

// GetVideoInfo get video info
func GetVideoInfo(videoID string, channelID string) (*VideoInfo, error) {
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%v", videoID)
	content, err := GetYouTubeHTML(url)
	if err != nil {
		return nil, err
	}

	var info VideoInfo

	doc, err := htmlquery.Parse(bytes.NewReader([]byte(content)))
	if err != nil {
		return nil, err
	}

	var one *html.Node

	// get video name
	one = htmlquery.FindOne(doc, `//meta[@itemprop="name"]/@content`)
	if one == nil {
		return nil, errors.New("get name failed")
	}
	info.Name = strings.TrimSpace(htmlquery.InnerText(one))

	// get videoID
	one = htmlquery.FindOne(doc, `//meta[@itemprop="identifier"]/@content`)
	if one == nil {
		return nil, errors.New("get videoID failed")
	}
	info.VideoID = strings.TrimSpace(htmlquery.InnerText(one))
	info.Link = fmt.Sprintf("https://www.youtube.com/watch?v=%v", info.VideoID)

	// channelID
	info.ChannelID = channelID

	// get datePublished
	one = htmlquery.FindOne(doc, `//meta[@itemprop="datePublished"]/@content`)
	if one == nil {
		return nil, errors.New("get datePublished failed")
	}
	info.DatePublished = strings.TrimSpace(htmlquery.InnerText(one))

	// get author name
	one = htmlquery.FindOne(doc, `//span[@itemprop="author"]/link[@itemprop="name"]/@content`)
	if one == nil {
		return nil, errors.New("get author name failed")
	}
	info.Author = strings.TrimSpace(htmlquery.InnerText(one))

	return &info, nil
}
