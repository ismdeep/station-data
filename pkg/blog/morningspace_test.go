package blog

import "testing"

func TestMorningSpace(t *testing.T) {
	testBlogger(t, &MorningSpace{})
}
