package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	Title  string `json: "title"`
	Author string `json:"author"`
}

var jsonString = []byte(`
[
	{"title": "The Art of Community", "author": "Jono Bacon"},
	{"title": "Mithril", "author": "Yoshiku Shibukawa"}
]`)

func main() {
	var books []Book
	err := json.Unmarshal(jsonString, &books)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		fmt.Println(book)
	}

	d, _ := json.Marshal(Book{"눈을 뜨자!", "Cody Lindely"})
	log.Println(string(d))
}
