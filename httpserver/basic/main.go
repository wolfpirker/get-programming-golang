package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// https://golang.org/pkg/net/http/
// https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
