package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		//쿠키가 있기 때문에 한 번 다녀간 적이 있다.
		fmt.Fprintf(w, "<html><body>다녀간 적 있음</body></html>\n")
	} else {
		fmt.Fprintf(w, "<html><body>첫방문</body></html>\n")
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening : 18888")
	httpServer.Addr = ":5000"
	log.Println(httpServer.ListenAndServe())
}
