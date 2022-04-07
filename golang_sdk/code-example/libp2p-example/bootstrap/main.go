package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("\n the error is : %v", err)
	}
}

func main() {
	host, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
	)
	CheckErr(err)
	fmt.Println(host.ID(), "   ", host.Addrs())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
	host.Close()

}
