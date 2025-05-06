package main

import (
	"fmt"
	"golang_journey/utils/link"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

// 全局变量，用于存储已访问的 URL
var visited = make(map[string]bool)
var mu sync.Mutex

func main() {

	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}

}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := link.ExtractLinks(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	return list
}

func printUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Println(string(b))
}
