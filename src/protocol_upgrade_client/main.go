package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	// TCP 소켓 개방
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)

	//요청을 작성해 소켓에 직접 써넣기
	request, _ := http.NewRequest("GET", "http://127.0.0.1:5000/upgrade", nil)
	request.Header.Set("Connection", "Upgrade")
	request.Header.Set("Upgrade", "MyProtocol")
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	// 소켓에서 데이터를 읽어와 응답 분석
	resp, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println("Status: ", resp.Status)
	log.Println("Status: ", resp.Header)

	// 오리지널 통신
	counter := 10
	for {
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Println("<-", string(bytes.TrimSpace(data)))
		fmt.Fprintf(conn, "%d\n", counter)
		fmt.Println("->", counter)
		counter--
	}
}
