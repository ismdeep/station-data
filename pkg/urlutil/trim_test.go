package urlutil

import "testing"

func TestTrimPath(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				uri: "https://test.lemuria.sh/",
			},
			want: "https://test.lemuria.sh/",
		},
		{
			name: "",
			args: args{
				uri: "https://test.lemuria.sh//",
			},
			want: "https://test.lemuria.sh/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimPath(tt.args.uri); got != tt.want {
				t.Errorf("TrimPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
