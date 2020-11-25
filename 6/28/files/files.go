package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("./5/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files, err = ioutil.ReadDir("/etc/hosts")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
