package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type URLCache struct {
	cachedURLs map[string]int
	mux sync.Mutex
}

func (cache URLCache) putIfAbsent (url string) bool{
	result := false
	cache.mux.Lock()
	if cache.cachedURLs[url]==1 {
		result = true
	} else {
		cache.cachedURLs[url]=1
	}
	defer cache.mux.Unlock()
	return result
}
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		//ch <- print(err)
		return
	}
	ch <- fmt.Sprintf("found: %s %q\n", url, body)
	for _, u := range urls {
		if !urlCache.putIfAbsent(u){
			go Crawl(u, depth-1, fetcher)
		}
	}
	return
}

func main() {
	go Crawl("http://golang.org/", 4, fetcher)
	for i:=1; i<=3; i++ {
		fmt.Print(<-ch);
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

var urlCache = URLCache{cachedURLs: make(map[string]int)}
var ch = make (chan string)
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
