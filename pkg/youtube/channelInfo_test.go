package youtube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChannelInfo(t *testing.T) {
	type args struct {
		channelHomeURL string
	}
	tests := []struct {
		name    string
		args    args
		want    *ChannelInfo
		wantErr bool
	}{
		{
			name: "",
			args: args{
				channelHomeURL: "https://www.youtube.com/@GOTO-",
			},
			want: &ChannelInfo{
				ChannelID: "UCs_tLP3AiwYKwdUHpltJPuA",
				Name:      "GOTO Conferences",
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				channelHomeURL: "https://www.youtube.com/@laogao",
			},
			want: &ChannelInfo{
				ChannelID: "UCMUnInmOkrWN4gof9KlhNmQ",
				Name:      "老高與小茉 Mr & Mrs Gao",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChannelInfo(tt.args.channelHomeURL)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
