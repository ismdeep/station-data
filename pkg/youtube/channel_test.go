package youtube

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllVideos(t *testing.T) {
	channelID := "UCMUnInmOkrWN4gof9KlhNmQ"
	videos, err := GetAllVideos(channelID)
	assert.NoError(t, err)
	var wg sync.WaitGroup
	for i, videoID := range videos {
		wg.Add(1)
		go func(idx int, videoID string) {
			defer wg.Done()
			// Get video info by id
			videoInfo, err := GetVideoInfo(videoID, channelID)
			if err == nil {
				t.Logf("videos[%v] = https://www.youtube.com/watch?v=%v %v", i, videoID, videoInfo.Name)
			}
		}(i, videoID)
	}
	wg.Wait()
}
