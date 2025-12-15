package search

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

// SearchPeople searches LinkedIn people and collects profile URLs
func SearchPeople(page *rod.Page, keyword string, limit int) []string {
	fmt.Println("Searching people for keyword:", keyword)

	searchURL := "https://www.linkedin.com/search/results/people/?keywords=" + keyword
	page.MustNavigate(searchURL)
	page.MustWaitLoad()

	seen := make(map[string]bool)
	results := []string{}

	for len(results) < limit {
		// Scroll to load more results
		page.Mouse.MustScroll(0, 1200)

		time.Sleep(3 * time.Second)

		links := page.MustElements("a")

		for _, link := range links {
			var url string

			func() {
				defer func() {
					recover()
				}()

				href := link.MustAttribute("href")
				url = *href
			}()

			if url == "" {
				continue
			}

			if strings.Contains(url, "/in/") && !seen[url] {
				seen[url] = true
				results = append(results, url)
				fmt.Println("Found profile:", url)

				if len(results) >= limit {
					break
				}
			}
		}
	}

	fmt.Println("Total profiles collected:", len(results))
	return results
}
