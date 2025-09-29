package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: 500 * time.Millisecond}
	req, err := c.Get("https://golang.org")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
