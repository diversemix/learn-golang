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

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(chDone chan bool, visted *VisitedUrls, url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	if visted.Get(url) {
		fmt.Println("...skipping:", url)
		chDone <- true
		return
	}

	if depth <= 0 {
		fmt.Println("Max depth reached !!!")
		chDone <- false
		return
	}
	body, urls, err := fetcher.Fetch(url)
	visted.Set(url)
	if err != nil {
		fmt.Println(err)
		chDone <- false
		return
	}

	fmt.Printf("found: %s %q urls=%d\n", url, body, len(urls))

	ch := make(chan bool)
	for _, u := range urls {
		go Crawl(ch, visted, u, depth-1, fetcher)
	}
	for range urls {
		<-ch
	}
	chDone <- true
	return
}

func main() {
	visted := VisitedUrls{urls: make(map[string]bool)}
	ch := make(chan bool)

	go Crawl(ch, &visted, "https://golang.org/", 4, fetcher)
	fmt.Println("MAIN... waiting...")
	fmt.Println("MAIN... done!", <-ch)
}

type VisitedUrls struct {
	urls map[string]bool
	mux  sync.Mutex
}

func (v *VisitedUrls) Set(key string) {
	v.mux.Lock()
	v.urls[key] = true
	v.mux.Unlock()
}

func (v *VisitedUrls) Get(key string) bool {
	v.mux.Lock()
	defer v.mux.Unlock()
	return v.urls[key]
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
