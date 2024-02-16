package bilibili

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSpace(t *testing.T) {
	type args struct {
		spaceID string
	}
	tests := []struct {
		name string
		args args
		want *Space
	}{
		{
			name: "",
			args: args{
				spaceID: "46708782",
			},
			want: &Space{
				ID: "46708782",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpace(tt.args.spaceID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSpaceByURL(t *testing.T) {
	type args struct {
		spaceURL string
	}
	tests := []struct {
		name string
		args args
		want *Space
	}{
		{
			name: "",
			args: args{
				spaceURL: "https://space.bilibili.com/46708782",
			},
			want: &Space{
				ID: "46708782",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpaceByURL(tt.args.spaceURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpaceByURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpace_GetInfo(t *testing.T) {
	type fields struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *SpaceInfo
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				ID: "46708782",
			},
			want: &SpaceInfo{
				ID:   "46708782",
				Name: "鲁大能",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := &Space{
				ID: tt.fields.ID,
			}
			got, err := receiver.GetInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpaceInfo_GetVideos(t *testing.T) {
	tests := []struct {
		name     string
		receiver Space
		wantErr  bool
	}{
		{
			name: "",
			receiver: Space{
				ID: "46708782",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.receiver.GetVideos()
			assert.Equal(t, tt.wantErr, err != nil)
			for i, v := range got {
				t.Logf("got video %v = %v %v", i, v.ID, v.Title)
			}
		})
	}
}
