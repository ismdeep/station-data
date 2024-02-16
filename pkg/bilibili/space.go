package bilibili

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Space model
type Space struct {
	ID string
}

// NewSpace new a space by id
func NewSpace(spaceID string) *Space {
	return &Space{ID: spaceID}
}

// NewSpaceByURL new a space with space home url
func NewSpaceByURL(spaceURL string) *Space {
	return NewSpace(spaceURL[strings.LastIndex(spaceURL, "/")+1:])
}

// GetInfo get space info
func (receiver *Space) GetInfo() (*SpaceInfo, error) {
	var info struct {
		MID  int64  `json:"mid"`
		Name string `json:"name"`
	}

	if err := getInfo(fmt.Sprintf("/x/space/acc/info?mid=%v", receiver.ID), &info); err != nil {
		return nil, err
	}

	if fmt.Sprintf("%v", info.MID) != receiver.ID {
		return nil, errors.New("invalid data")
	}

	return &SpaceInfo{
		ID:   fmt.Sprintf("%v", info.MID),
		Name: info.Name,
	}, nil
}

// GetVideos get space all videos
func (receiver *Space) GetVideos() ([]Video, error) {
	var lst []Video

	// get video count
	cnt, err := receiver.GetVideoCount()
	if err != nil {
		return nil, err
	}

	// get videos
	pageSize := 30
	for pageNo := 1; pageNo <= (cnt+1)/pageSize; pageNo++ {
		uri := fmt.Sprintf("/x/space/wbi/arc/search?mid=%v&ps=%v&tid=0&pn=%v&keyword=&order=pubdate&order_avoided=true", receiver.ID, pageSize, pageNo)
		var data struct {
			List struct {
				VList []struct {
					BVID        string `json:"bvid"`
					Title       string `json:"title"`
					Description string `json:"description"`
					Created     int64  `json:"created"`
				} `json:"vlist"`
			} `json:"list"`
		}
		if err := getInfo(uri, &data); err != nil {
			return nil, err
		}
		for _, item := range data.List.VList {
			lst = append(lst, Video{
				ID:          item.BVID,
				Title:       item.Title,
				Description: item.Description,
				Created:     time.Unix(item.Created, 0),
			})
		}
	}

	return lst, nil
}

// GetVideoCount get video count
func (receiver *Space) GetVideoCount() (int, error) {
	uri := fmt.Sprintf("/x/space/navnum?mid=%v", receiver.ID)
	var data struct {
		Video int `json:"video"`
	}

	if err := getInfo(uri, &data); err != nil {
		return 0, err
	}

	return data.Video, nil
}
