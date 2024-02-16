package response

import "time"

// Monitor struct
type Monitor struct {
	Name        string    `json:"name"`
	Msg         string    `json:"msg"`
	Time        time.Time `json:"time"`
	Status      int       `json:"status"`
	StatusText  string    `json:"status_text"`
	StatusColor string    `json:"status_color"`
}
