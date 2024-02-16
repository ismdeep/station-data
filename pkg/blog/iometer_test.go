package blog

import "testing"

func TestIOMeter(t *testing.T) {
	testBlogger(t, &IOMeter{})
}
