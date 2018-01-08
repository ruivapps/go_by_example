package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func HttpGet(url string, ch chan<- string, count int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	//channel is use for go routine communicate
	ch <- fmt.Sprintf("%d %d", count, len(bytes))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("http_concurrent_get url url url ...\n")
		os.Exit(1)
	}
	// make a channel that hold string (you can make channel for int, boolean etc)
	ch := make(chan string)
	begin := time.Now()
	// range will return index, value, so we get 0 url1, 1 url2 ...
	for i, url := range os.Args[1:] {
		go HttpGet(url, ch, i)
	}
	// this is just loop without request any value
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Println("time spent:", time.Since(begin).Seconds())
}
