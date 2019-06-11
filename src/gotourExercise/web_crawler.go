package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, m *sync.Map, ch chan []string) {
	defer wg.Done()
	done, _ := m.Load(url)
	if done == true || depth <= 0 {
		return
	}
	m.Store(url, true)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- []string{url, body}
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, m, ch)
	}
}

func main() {
	m := sync.Map{}
	ch := make(chan []string)
	wg.Add(1)
	go Crawl("http://golang.org/", 4, fetcher, &m, ch)
	go func() {
		wg.Wait()
		close(ch)
	}()
	for res := range ch {
		fmt.Printf("found: %s %q\n", res[0], res[1])
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
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
