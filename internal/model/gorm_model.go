package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/ismdeep/station-data/pkg/quantumid"
)

// Blog model
type Blog struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Source      string `gorm:"type:varchar(255);index;"`
	Link        string `gorm:"type:varchar(768);index;"`
	Title       string `gorm:"type:varchar(768);index;"`
	Author      string `gorm:"type:varchar(255);index;"`
	PublishedAt time.Time
	Content     string
	ReadTime    *time.Time
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ConfGithubRepo model
type ConfGithubRepo struct {
	Owner string `gorm:"type:varchar(255);primaryKey" json:"owner"`
	Repo  string `gorm:"type:varchar(255);primaryKey" json:"repo"`
}

// ConfYoutubeChannel model
type ConfYoutubeChannel struct {
	ID   string `gorm:"type:varchar(50);primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}

// ConfYoutubePlaylist YouTube playlist model
type ConfYoutubePlaylist struct {
	ID   string `gorm:"type:varchar(128);primaryKey" json:"id"`
	Name string `gorm:"type:varchar(1024)" json:"name"`
}

// ConfBilibiliSpace model
type ConfBilibiliSpace struct {
	ID   string `gorm:"type:varchar(50);primaryKey" json:"id"`
	Name string `gorm:"type:varchar(1024)" json:"name"`
}

// User model
type User struct {
	ID        string `gorm:"type:varchar(50);primaryKey"`
	Username  string `gorm:"type:varchar(50);unique"`
	Digest    string `gorm:"type:varchar(255)"`
	CreatedAt int64  `gorm:"type:bigint;not null;autoUpdateTime:false"`
	UpdatedAt int64  `gorm:"type:bigint;not null;autoUpdateTime:false"`
}

// BeforeCreate before create
func (receiver *User) BeforeCreate(tx *gorm.DB) error {
	receiver.ID = quantumid.NewString()
	now := time.Now().UnixNano()
	receiver.CreatedAt = now
	receiver.UpdatedAt = now
	return nil
}
