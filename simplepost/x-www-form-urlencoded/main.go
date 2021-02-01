package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm("http://127.0.0.1:5000", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}
