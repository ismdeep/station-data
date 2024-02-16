package blog

import "testing"

func TestEltonMinetto(t *testing.T) {
	testBlogger(t, &EltonMinetto{})
}
