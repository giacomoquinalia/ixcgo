package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// START OMIT
func encriptMyData(key, data string) {
	var encripted []byte
	block, _ := aes.NewCipher([]byte(key))
	cipher.NewGCM(block)
	block.Encrypt([]byte(data), encripted)
	fmt.Printf("Encrypted data: %s\n", encripted)
}

// END OMIT

func main() {
	encriptMyData("my key", "my data")
}
