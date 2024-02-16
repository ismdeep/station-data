package sha256util

import "testing"

func TestSHA256(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				content: []byte("hello world"),
			},
			want: "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256(tt.args.content); got != tt.want {
				t.Errorf("SHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}
