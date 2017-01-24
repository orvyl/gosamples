package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix("http://", url) {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Something went wrong in fetching the data: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Printf("Something went wrong in reading the data: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n\n HTTP status is: %v", b, resp.Status)
	}
}
