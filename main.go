package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var url string = os.Args[1]

	r, err := getPageContent(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
}

func getPageContent(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
