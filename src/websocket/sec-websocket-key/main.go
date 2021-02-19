package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	clientSecretSrc := "the sample nonce"
	key := base64.StdEncoding.EncodeToString([]byte(clientSecretSrc))
	fmt.Printf("Sec-WebSocket-Key: %s\n", key)

	salt := "ASDWEABASO-E918-AMEZ-AIUPBWIUBK1"
	hash := sha1.Sum([]byte(key + salt))
	accept := base64.StdEncoding.EncodeToString(hash[:])
	fmt.Printf("Sec-WebSocket-Accept: %s\n", accept)
}
