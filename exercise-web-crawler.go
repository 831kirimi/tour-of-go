package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	
	//fmt.Printf("depth = %v\n", depth)
	map1 := make(map[string]bool)
	ch := make(chan string,100)
	var prev int
	var mu sync.Mutex
	
	var crawl func(string, int)
	
    crawl = func(url string, depth int) {
	if depth <= 0 {
		return
	}
		ch <- url
	body, urls, err := fetcher.Fetch(url)
	if map1[url] {return}
	mu.Lock()
    
	if err != nil {
		//fmt.Printf("error,depth=%v\n",depth)
		fmt.Println(err)
		map1[url] = true
		return
	}
	
	fmt.Printf("found: %s %q\n", url, body)
	map1[url] = true
	for _, u := range urls {
		
		go crawl(u, depth-1)
		
	}
		defer mu.Unlock()
	}
	crawl(url, depth)
	for {
        if len(ch) != prev {
            prev = len(ch)
        } else {
            break
        }
        time.Sleep(time.Millisecond)
    }
	return
}

func main() {
	
	Crawl("https://golang.org/", 4, fetcher)
	
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

