package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	part := make(textproto.MIMEHeader)                                                  // fileWriter, err := writer.CreateFormFile("tumbnail", "photo.jpg")
	part.Set("Content-Type", "image/jpeg")                                              // if err != nil {
	part.Set("Content-Disposition", `form-data; name="thumbnail; filename="photo.jpg"`) // 	panic(err)
	fileWriter, err := writer.CreatePart(part)                                          // }
	if err != nil {                                                                     // readFile, err := os.Open("photo.jpg")
		panic(err) // if err != nil {
	} // 	panic(err)
	readFile, err := os.Open("photo.jpg") // }
	if err != nil {                       // defer readFile.Close()
		panic(err) // io.Copy(fileWriter, readFile)
	}
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://127.0.0.1:5000", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status", resp.Status)
}
