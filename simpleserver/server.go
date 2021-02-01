package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/k0kubun/pp"
)

func handlerDigest(w http.ResponseWriter, r *http.Request) {
	pp.Printf("URL: %s\n", r.URL.String())
	pp.Printf("Query: %v\n", r.URL.Query())
	pp.Printf("Proto: %s\n", r.Proto)
	pp.Printf("Method: %s\n", r.Method)
	pp.Printf("Header: %v\n", r.Header)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("--body--\n%s\n", string(body))
	if _, ok := r.Header["Authorization"]; !ok {
		w.Header().Add("www-Authenticate", `Digest realm="Secret Zone", 
											nonce="TgLc25U2BQA=f510a27800473e121489308afwe83248c",
											algorithm=MD5, qop="auth`)
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		fmt.Fprintf(w, "<html><body>secret page</body></html>\n")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		//쿠키가 있기 때문에 한 번 다녀간 적이 있다.
		fmt.Print(r.Header["Cookie"])
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

	http.HandleFunc("/cookie", handler)
	http.HandleFunc("/digest", handlerDigest)

	log.Println("start http listening : 18888")
	httpServer.Addr = ":5000"
	log.Println(httpServer.ListenAndServe())
}
