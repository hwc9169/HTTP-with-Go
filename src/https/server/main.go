package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Start http listening :5050")
	err := http.ListenAndServeTLS(":5050", "/home/hwc9169/https/server.crt", "/home/hwc9169/https/server.key", nil)
	log.Println(err)
}
