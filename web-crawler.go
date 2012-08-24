package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	const (
		NOT_VISITED = 1
		VISITED     = 2
	)
	urls := make(map[string]int)

	urls[url] = NOT_VISITED
	for i := 0; i < depth; i++ {
		nr := 0
		ch := make(chan []string)
		for k, v := range urls {
			if v == NOT_VISITED {
				nr++
				go func(inputUrl string) {
					body, urls, err := fetcher.Fetch(inputUrl)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("found: %s %q\n", inputUrl, body)
					}
					// we must send something back, even it's nil
					ch <- urls
				}(k)
				urls[k] = VISITED
			}
		}
		for j := 0; j < nr; j++ {
			l := <-ch
			for _, ll := range l {
				if urls[ll] != VISITED {
					urls[ll] = NOT_VISITED
				}
			}
		}
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
