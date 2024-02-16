package archiver

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ismdeep/station-data/pkg/blog"
	"github.com/ismdeep/station-data/pkg/sha256util"
	"github.com/ismdeep/station-data/pkg/urlutil"
)

func WriteMeta(source string) error {
	_, err := os.Stat(fmt.Sprintf("./data/%v/meta.json", source))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if err == nil {
		return nil
	}

	type MetaInfo struct {
		Source string `json:"source"`
	}

	metaInfo := MetaInfo{
		Source: source,
	}

	raw, err := json.MarshalIndent(metaInfo, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(fmt.Sprintf("./data/%v/meta.json", source), raw, 0640); err != nil {
		return err
	}

	return nil
}

func LoadArchived(source string) (map[string]string, error) {
	content, err := os.ReadFile(fmt.Sprintf("./data/%v/index.txt", source))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	if errors.Is(err, os.ErrNotExist) {
		return make(map[string]string), nil
	}

	results := make(map[string]string)

	lst := strings.Split(string(content), "\n")
	for _, s := range lst {
		s = strings.Trim(s, "\r\n")
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		sha256 := s[:strings.Index(s, " ")]
		uri := s[strings.Index(s, " ")+1:]
		results[uri] = sha256
	}

	return results, nil
}

func ProcessPages(blogger blog.BloggerInterface, archivedLst map[string]string, pageUrls []string) error {
	indexFile, err := os.OpenFile(fmt.Sprintf("./data/%v/index.txt", blogger.GetBloggerName()), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0640)
	if err != nil {
		return err
	}

	pageUrlsChan := make(chan string, 1024)
	go func() {
		for _, pageUrl := range pageUrls {
			pageUrlsChan <- pageUrl
		}
		close(pageUrlsChan)
	}()

	var errLst []error

	var wg sync.WaitGroup
	for range 16 {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()

			for pageUrl := range pageUrlsChan {
				links, err := blogger.GetLinksFromPage(pageUrl)
				if err != nil {
					errLst = append(errLst, err)
					continue
				}
				fmt.Println("OK:", blogger.GetBloggerName(), "blog links", pageUrl)
				for _, link := range links {
					link = urlutil.TrimPath(link)
					if _, ok := archivedLst[link]; ok {
						continue
					}
					blogInfo, err := blogger.GetBlogInfo(link)
					if err != nil {
						errLst = append(errLst, err)
						continue
					}

					raw, err := json.MarshalIndent(blogInfo, "", "  ")
					if err != nil {
						errLst = append(errLst, err)
						continue
					}
					if err := os.WriteFile(fmt.Sprintf("./data/%v/data/%v.json", blogger.GetBloggerName(), sha256util.SHA256([]byte(link))), raw, 0640); err != nil {
						errLst = append(errLst, err)
						continue
					}

					if _, err := indexFile.WriteString(fmt.Sprintf("%v %v\n", sha256util.SHA256([]byte(link)), link)); err != nil {
						errLst = append(errLst, err)
						continue
					}

					fmt.Println("OK:", blogger.GetBloggerName(), "blog info ", link)
				}
			}
		}()
	}
	wg.Wait()

	return errors.Join(errLst...)
}

func Archive(blogger blog.BloggerInterface) error {
	// mkdir
	if err := os.MkdirAll(fmt.Sprintf("./data/%v/data", blogger.GetBloggerName()), 0750); err != nil {
		return err
	}
	fmt.Println("OK:", blogger.GetBloggerName(), "make directory")

	// write meta.json
	if err := WriteMeta(blogger.GetBloggerName()); err != nil {
		return err
	}
	fmt.Println("OK:", blogger.GetBloggerName(), "write meta.json")

	// load archived links
	archivedLst, err := LoadArchived(blogger.GetBloggerName())
	if err != nil {
		return err
	}
	fmt.Println("OK:", blogger.GetBloggerName(), "load archived list")

	// page url
	var pageUrls []string
	switch os.Getenv("STATION_ARCHIVER_ALL_PAGE") {
	case "true":
		pageUrls = blogger.GetAllPageURLs()
	default:
		pageUrls = []string{
			blogger.GetFirstPageURL(),
		}
	}

	if err := ProcessPages(blogger, archivedLst, pageUrls); err != nil {
		return err
	}
	fmt.Println("OK:", blogger.GetBloggerName(), "finished")

	return nil
}
