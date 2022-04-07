package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fp, fperr := os.OpenFile("input.jpg", os.O_RDONLY, 0666)
	if fperr != nil {
		fmt.Println("\n Error opening file")
	}
	//RSA Public and Private Key Generation
	PrivateKey, gen_err := rsa.GenerateKey(rand.Reader, 512)
	if gen_err != nil {
		log.Panic("error on generating keys")
	}
	// fmt.Printf("\n Private Key %v and Public Key %v", *PrivateKey, PrivateKey.PublicKey)

	var image_data = new(bytes.Buffer)

	io.Copy(image_data, fp)

	fmt.Printf("\n Image size %v\n", len(image_data.Bytes()))

	output_bytes, enc_err := rsa.EncryptPKCS1v15(rand.Reader, &PrivateKey.PublicKey, image_data.Bytes()[:50])
	if enc_err != nil {
		log.Panic(enc_err)
	}
	fmt.Println(output_bytes)

	fp.Close()
}
