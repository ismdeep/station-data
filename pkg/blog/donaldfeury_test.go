package blog

import "testing"

func TestDonaldFeury(t *testing.T) {
	testBlogger(t, &DonaldFeury{})
}
