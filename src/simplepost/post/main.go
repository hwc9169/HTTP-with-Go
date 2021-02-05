package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	// 파일을 post로 전달
	resp, err := http.Post("http://127.0.0.1:5000", "test/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)

	//"문자열"을 post 전달
	var str string = "텍스트"
	var reader *strings.Reader = strings.NewReader(str)
	resp, err = http.Post("http://127.0.0.1:5000", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)

}
