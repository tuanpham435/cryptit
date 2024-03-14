package main

import (
	"fmt"
	"github.com/tuanpham435/cryptit/decrypt"
	"github.com/tuanpham435/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Tuan")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
