package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
	"unsafe"
)

//defining a struct
type mystruct struct {
	x int
	y int
}

//defining a struct
type userdata struct {
	Name string `json:"Name"` //the Caps first one says that it is public
	Id   int    `json:"Id"`
}

//defining a struct for linked list
type Block struct {
	prev *Block
	data int
	next *Block
}

//define a function
func fun1(x int) int {
	return x
}

//define a method
func (x mystruct) method() int {
	return x.x + x.y
}

//interface
type myinterface interface {
	method() int
}

func Check_error(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	/******************************net         package**************************************/
	sm := http.ServeMux{}
	sm.ServeHTTP()

	sm.HandleFunc("/bipul", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nthis is bipul")
	})

	sm.HandleFunc("/bipul", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nthis is bipul")
	})

	server := http.Server{Addr: "localhost:5678", Handler: &sm}
	fmt.Println(server.ReadTimeout)
	server.ListenAndServe()
	return
	handler := http.RedirectHandler("/time", 200)
	HttpServer := &http.Server{Addr: "localhost:7890", Handler: handler}
	http_err := HttpServer.ListenAndServe()
	if http_err != nil {
		log.Panic(http_err)
	}
	addr, resolve_err := net.ResolveIPAddr("ip", "google.com")
	if resolve_err != nil {
		log.Panic(resolve_err)
	}
	fmt.Printf("\n Ip addr is %v", addr)

	addresses, lookup_err := net.LookupHost("google.com")
	if lookup_err != nil {
		log.Panic(lookup_err)
	}
	fmt.Printf("\n Ip addr is %v", addresses)

	tcpAddr, service_err := net.ResolveTCPAddr("tcp4", "localhost:54783")
	if service_err != nil {
		log.Panic(service_err)
	}
	fmt.Printf("\n Ip addr is %v", tcpAddr)

	connection, dialup_err := net.DialTCP("tcp4", nil, tcpAddr)
	if dialup_err != nil {
		log.Panic(dialup_err)
	}
	// nc -l --local-port=54783 do this before this
	_, write_err := connection.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if write_err != nil {
		log.Panic(write_err)
	}

	fmt.Printf("\n typee %T", connection.Write)
	/***************************************************************************************/
	/*******************************Input Output Package************************************/
	file1, file1_err := os.Create("file1.txt")
	file2, file2_err := os.Open("testfile.txt")
	if file2_err != nil {
		log.Panic(file2_err)
	}
	if file1_err != nil {
		log.Panic(file1_err)
	}

	io.Copy(file1, file2)

	file1.Close()
	file2.Close()
	/***************************************************************************************/

	/******************************Data Serialization and Deserialization*******************/
	struct_data := &userdata{Name: "bipul", Id: 1}
	json_data, json_encoder_err := json.Marshal(struct_data)
	if json_encoder_err != nil {
		log.Panic(json_encoder_err)
	}
	fmt.Printf(string(json_data)) //json_data is in []byte

	/***************************************************************************************/
	//variable with specified type
	var x int = 2
	fmt.Printf("value = %X and type = %T", x, x)

	//make a struct type, semicolon or new line to break
	var y mystruct = mystruct{x: 1, y: 2}
	fmt.Printf("\nvalue = %X and type = %T", y, y)

	//call a function
	var result int = fun1(2)
	fmt.Printf("\nvalue = %X and type = %T", result, result)

	//calling interface method
	var caller myinterface = y //casted struct to interface
	var interface_result int = caller.method()
	fmt.Printf("\nresult = %X", interface_result)

	//building a linked list from Blocks
	var b1 Block = Block{prev: nil, data: 0, next: nil}
	var b2 Block = Block{prev: &b1, data: 1, next: nil}
	b1.next = &b2
	fmt.Printf("\nb1 = %X, b2 = %X", b1, *b1.next)

	//channel
	ch := make(chan string)
	go func() {
		var p = <-ch
		fmt.Printf("\np = %v", p)
	}()
	ch <- "paper"
	close(ch)

	//goroutine
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Printf("\nthis is first  routine")
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Printf("\nthis is second routine")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 5)
	fmt.Printf("\n who is running first?")

	//hashing in go
	var h = sha1.New()
	var input_string string = "hello! this is Bipul"
	var input_in_bytes []byte = []byte(input_string)
	h.Write(input_in_bytes)
	var hash = h.Sum(nil)
	fmt.Printf("\n hash value = %X\n", hash)

	//type casting
	var s = int(5)
	fmt.Println(s)

	//size of
	fmt.Printf("\nSize of %v is %x", 1, unsafe.Sizeof(1))

	//Buffer
	var buff = new(bytes.Buffer) //new() only returns pointers to initialized memory.
	var err = binary.Write(buff, binary.BigEndian, uint64(257))
	if err != nil {
		log.Panic("\nerror")
	}
	fmt.Printf("\n bytes are %v", buff.Bytes())

	//commputer stores number but the same number can be represented in many ways. at the storage level it stores number in 2's complement.

	//handling numbers more than uint64
	var biginteger big.Int
	biginteger.SetBytes([]byte{123})
	fmt.Printf("\n biginteger = %v", biginteger)

	/****************************************SECURITY****************************************/
	//sha-256 hashing with codition such that the hashed value is smaller than a specified number
	var target = big.NewInt(1)
	target = target.Lsh(target, 256-6) //that means we want first 6 bits to be 0's
	var message = "HI I am bipul"
	var calculated_hash big.Int
	for i := 0; i < math.MaxInt64; i++ {
		hashed := sha256.Sum256([]byte(message + string(i)))
		calculated_hash.SetBytes(hashed[:])
		if calculated_hash.Cmp(target) == -1 { //so that calculated_hash have also 6 0's at first
			fmt.Printf("\nfinal hash %v", calculated_hash)
			break
		}
	}

	//RSA Public and Private Key Generation
	PrivateKey, gen_err := rsa.GenerateKey(rand.Reader, 512)
	if gen_err != nil {
		log.Panic("error on generating keys")
	}
	fmt.Printf("\n Private Key %v and Public Key %v", *PrivateKey, PrivateKey.PublicKey)

	//signing a hash using RSA using private key
	hashed := sha256.Sum256([]byte("HI this is a data but this is too long "))
	signature, signature_err := rsa.SignPKCS1v15(rand.Reader, PrivateKey, crypto.SHA256, hashed[:])
	if signature_err != nil {
		panic(signature_err)
	}
	fmt.Printf("\n Signature is %v the length of hashed is %v", signature, len(hashed))

	// verification of above signature using public key
	verified_err := rsa.VerifyPKCS1v15(&PrivateKey.PublicKey, crypto.SHA256, hashed[:], signature[:])
	if verified_err != nil {
		panic(verified_err)
	}
	fmt.Printf("\n SUCCESS")
	/***************************************************************************************/

}

// https://pkg.go.dev/std
