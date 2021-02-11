package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	// 인증서를 읽어 드림
	cert, err := ioutil.ReadFile("/home/hwc9169/https/ca.crt")
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}
	tlsConfig.BuildNameToCertificate()

	// 클라이언트 작성
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	// 통신한다
	resp, err := client.Get("https://localhost:5050")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
