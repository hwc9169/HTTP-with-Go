package main

import(
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

var html []byte

func handlerHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write(html)
}


func handlerPrimeSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	closeNotify := w.(http.CloseNotifier).CloseNotify()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origni", "*")

	var num int64 = 1
	for id :=1; id <= 100; id++ {
		// 통신이 끊겨도 종료
		select {
		case <- closeNotify:
			fmt.Println("Connection closed from client")
			return
		default:
			// do nothing
		}
		for {
			num++
			if big.NewInt(num).ProbablyPrime(20) {
				fmt.Println(num)
				fmt.Fprintf(w, "id: %d\n", id)
				fmt.Fprintf(w, "event: %s\n", "prime")
				fmt.Fprintf(w, "data: %d\n\n", num)
				flusher.Flush()
				time.Sleep(time.Second)
				break
			}
		}
		time.Sleep(time.Second)
	}
	fmt.Println("Connection closed from server")
}

func main() {
	var err error
	html, err = ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handlerHtml)
	http.HandleFunc("/prime", handlerPrimeSSE)
	fmt.Println("start http listening :18888")
	err = http.ListenAndServe(":18888", nil)
	fmt.Println(err)
}