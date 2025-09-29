package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Go"}`))
	req, err := c.Post("https://golang.org", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	io.CopyBuffer(os.Stdout, req.Body, nil)
}
