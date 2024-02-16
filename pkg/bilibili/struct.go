package bilibili

import "time"

// SpaceInfo space info
type SpaceInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Video info
type Video struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
}
