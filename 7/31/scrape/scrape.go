package main

import (
	"fmt"
	"sync"
	"time"
)

// Visited tracks whether web pages have been visited.
// Its methods may be used concurrently with one another.
type Visited struct {
	// mu guards the visited map.
	mu      sync.Mutex
	visited map[string]int
}

// VisitLink tracks that the page with the given URL has
// been visited, and returns the updated link count.
func (v *Visited) VisitLink(url string) int {
	fmt.Printf("\n👀 Checking if %v has been visited", url)
	v.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	count := v.visited[url]
	count++
	v.visited[url] = count
	defer v.mu.Unlock()
	return count
}

func main() {
	c := Visited{visited: make(map[string]int)}

	for _, value := range []string{"www.google.com", "www.bing.com", "www.github.com"} {
		go c.VisitLink(value)
		time.Sleep(time.Second)
	}

}
