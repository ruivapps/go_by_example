/*
http post: https://golang.org/pkg/net/http/#Client.Post
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url string, data string) []byte {
	resp, err := http.Post(url, "application/json", strings.NewReader(data))
	/*
		use defer here instead calling resp.Body.Close() at end
		you can write defer anywhere in the function. when function end, it will run defer
		if you have stack defer, when function close, defer will be called in last in first out order (LIFO)
	*/
	defer resp.Body.Close()
	if err != nil {
		return []byte(err.Error())
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(err.Error())
	}
	//resp.Body.Close()
	return bytes
}

func main() {
	response := HttpPost("http://127.0.0.1:8080/fakeapi", "hello from golang client")
	if len(response) == 0 {
		fmt.Println("no data returned from server")
	} else {
		fmt.Println("server returns:\n" + string(response))
	}
}
