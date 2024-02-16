package release

import (
	"context"
	"fmt"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"

	"github.com/ismdeep/station-data/internal/schema"
)

// GetList get GitHub release list
func GetList(owner string, repo string, personalAccessToken string) ([]schema.Blog, error) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: personalAccessToken})
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)
	releases, _, err := client.Repositories.ListReleases(context.Background(), owner, repo, nil)
	if err != nil {
		return nil, err
	}

	var blogs []schema.Blog
	for i := len(releases) - 1; i >= 0; i-- {
		release := releases[i]
		var blog schema.Blog
		blog.Source = fmt.Sprintf("github.%v.%v.releases", owner, repo)
		if release.TagName == nil || release.HTMLURL == nil || release.PublishedAt == nil || release.Author == nil {
			continue
		}
		blog.Title = *release.TagName
		blog.Link = *release.HTMLURL
		blog.Date = release.PublishedAt.UTC()
		blog.Author = (*release.Author).GetLogin()

		blogs = append(blogs, blog)
	}

	return blogs, nil
}
