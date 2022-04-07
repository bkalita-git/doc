package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
	peerstore "github.com/libp2p/go-libp2p-core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
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
	if len(os.Args) > 1 {
		addr, err := multiaddr.NewMultiaddr(os.Args[1])
		CheckErr(err)
		fmt.Println(addr)
		peer, err := peerstore.AddrInfoFromP2pAddr(addr)
		CheckErr(err)
		conn_err := host.Connect(context.Background(), *peer)
		CheckErr(conn_err)

	} else {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		fmt.Println("Received signal, shutting down...")
	}
	host.Close()

}
