package blog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testBlogger(t *testing.T, blogger BloggerInterface) {
	assert.NotEqual(t, "", blogger.GetBloggerName())
	t.Logf("blogger: %v", blogger.GetBloggerName())

	assert.NotEqual(t, "", blogger.GetFirstPageURL())
	t.Logf("first page: %v", blogger.GetFirstPageURL())

	pageLinks := blogger.GetAllPageURLs()
	assert.Greater(t, len(pageLinks), 0)
	for i, link := range pageLinks {
		t.Logf("pageLinks[%v]: %v", i, link)
	}

	blogLinks, err := blogger.GetLinksFromPage(blogger.GetFirstPageURL())
	assert.NoError(t, err)
	assert.Greater(t, len(blogLinks), 0)
	for i, link := range blogLinks {
		t.Logf("blogLinks[%v]: %v", i, link)
	}

	blogInfo, err := blogger.GetBlogInfo(blogLinks[0])
	assert.NoError(t, err)

	t.Logf("blogInfo.Title = %v", blogInfo.Title)
	t.Logf("blogInfo.Author = %v", blogInfo.Author)
	t.Logf("blogInfo.Date = %v", blogInfo.Date)
	t.Logf("blogInfo.Content.length = %v", len(blogInfo.Content))
}

func TestMain(m *testing.M) {
	m.Run()
}
