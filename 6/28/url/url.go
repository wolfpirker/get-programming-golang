package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	urlv, err := url.Parse("https://ab .com")
	if err != nil{
		fmt.Println(err)
		fmt.Printf("error with url %q error:\n%#v", urlv, err)

		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op: ", e.Op)
			fmt.Println("URL: ", e.URL)
			fmt.Println("Err: ", e.Err)
		}
		os.Exit(1)
	} else{
		fmt.Printf("url %q ok!\n", urlv)
	
	}
	

}
