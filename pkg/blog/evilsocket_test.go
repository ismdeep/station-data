package blog

import "testing"

func TestEvilsocket(t *testing.T) {
	testBlogger(t, &Evilsocket{})
}
