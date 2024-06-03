package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// START OMIT
func hashMyData(data string) {
	sha256Hash := sha256.Sum256([]byte(data))
	sha512Hash := sha512.Sum512([]byte(data))
	md5Hash := md5.Sum([]byte(data))

	fmt.Printf("SHA-256: %x\n", sha256Hash)
	fmt.Printf("SHA-512: %x\n", sha512Hash)
	fmt.Printf("MD5: %x\n", md5Hash)
}

// END OMIT

func main() {
	hashMyData("my data")
}
