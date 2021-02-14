package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/k0kubun/pp"
)

func handlerUpgrade(w http.ResponseWriter, r *http.Request) {
	// 이 엔드포인트에서는 변경만 받아들인다.
	if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != "MyProtocol" {
		w.WriteHeader(400)
		return
	}
	fmt.Println("Upgrade to MyProtocol")

	// 소켓 획득
	hijacker := w.(http.Hijacker)
	conn, readWriter, err := hijacker.Hijack()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//  프로토콜이 변경 되었음을 응답
	response := http.Response{
		StatusCode: 101,
		Header:     make(http.Header),
	}
	response.Header.Set("Upgrade", "Myprotocol")
	response.Header.Set("Connection", "Upgrade")
	response.Write(conn)

	// 오리지널 통신 시작
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(readWriter, "%d\n", i)
		fmt.Println("->", i)
		readWriter.Flush() // Trigger "chunked" encoding and send a chunk
		recv, err := readWriter.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("<- %s", string(recv))
		time.Sleep(500 * time.Millisecond)
	}
}

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
	http.HandleFunc("/upgrade", handlerUpgrade)
	log.Println("start http listening : 5000")
	httpServer.Addr = ":5000"
	log.Println(httpServer.ListenAndServe())
}
