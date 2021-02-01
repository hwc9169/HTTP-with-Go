package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	http.DefaultClient = &http.Client{
		Jar: jar,
	}

	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:5000/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}

		log.Println(string(dump))
	}
}
