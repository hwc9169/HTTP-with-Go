package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Event struct {
	Name string
	ID string 
	Data string 
}


func EventSource(url string) (chan Event, context.Context, error){
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithCancel(req.Context())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode != 200 {
		return nil, nil, fmt.Errorf("Response Status Code should be 200, but %d\n",
	res.StatusCode)
	}
	events := make(chan Event)
	go receiveSSE(events, cancel, res)
	return events, ctx, nil
}


func receiveSSE(events chan Event, cancel context.CancelFunc, res *http.Response) {
	reader := bufio.NewReader(res.Body)
	event := Event{}
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			close(events)
			if err == io.EOF {
				cancel()
				return
			}
			panic(err)
		}
		fmt.Println("abc", line)
		switch {
		case bytes.HasPrefix(line, []byte(":ok")):
			//skip
		case bytes.HasPrefix(line, []byte("id:")):
			event.ID = string(line[4: len(line)-1])
		case bytes.HasPrefix(line, []byte("event:")):
			event.Name = string(line[7: len(line)-1])
		case bytes.HasPrefix(line, []byte("data:")):
			event.Data = string(line[6: len(line)-1])
		case bytes.Equal(line, []byte("\n")):
			if event.Data != "" {
				events <- event
			}
			event = Event{}
		default:
			fmt.Fprintf(os.Stderr, "Parse Error: %s\n", line)
			cancel()
			close(events)
		}
	}
}


func main() {
	events, ctx, err := EventSource("http://127.0.0.1:18888/prime")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case <- ctx.Done():
			return
		case event := <- events:

			fmt.Printf("Event(Id=%s, Event=%s): %s\n", event.ID, event.Name, event.Data)
		}
	}
}