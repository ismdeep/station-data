package blog

import (
	"testing"
)

func TestLogRocket(t *testing.T) {
	testBlogger(t, &LogRocket{})
}
