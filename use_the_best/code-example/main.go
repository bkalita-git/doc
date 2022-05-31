package main

import (
	"fmt"
	"time"
	"os"
)
func main(){
	//fmt.Println("Hello world!")
	//fp,_ := os.Open("test.txt")
	time.Sleep(time.Second*10)
	//var x uint32
	fp,_ := os.Open("test.txt")
	fmt.Println(fp)
	//time.Sleep(time.Minute*10)
	
}
