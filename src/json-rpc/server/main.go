package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 메서드가 속한 구조체
type Calculator int

// RPC로 외부에서 호출되는 메서드
func (c *Calculator) Multiply(args Args, result *int) error {
	log.Printf("Multiply called: %d, %d\n", args.A, args.B)
	*result = args.A * args.B
	return nil
}

// 외부에서 호출될 때의 인수
type Args struct {
	A, B int
}

func main() {
	calculator := new(Calculator)
	server := rpc.NewServer()
	server.Register(calculator)
	http.Handle(rpc.DefaultRPCPath, server)
	log.Println("start http listening: 5002")
	listener, err := net.Listen("tcp", ":5002")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
