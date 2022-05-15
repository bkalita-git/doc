package main

import (
	"context"
	"fmt"
	"go-p2p/communication"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
)

type nodeServer struct {
	communication.UnimplementedNodeServer
}

func (ns nodeServer) InitNode(nin communication.Node_InitNodeServer) error {
	fmt.Println(nin.Recv())
	fmt.Println(peer.FromContext(nin.Context()))
	nin.Send(&communication.NodeRes{Message: "I am server"})
	return nil
}

func reusePort(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(descriptor uintptr) {
		syscall.SetsockoptInt(int(descriptor), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		syscall.SetsockoptInt(int(descriptor), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
	})
}
func Server(wg *sync.WaitGroup, port_chan chan int) {
	config := &net.ListenConfig{Control: reusePort}
	listener, _ := config.Listen(context.Background(), "tcp4", "0.0.0.0:")
	fmt.Println("\nListening on port ", listener.Addr().String())
	addresses := strings.Split(listener.Addr().String(), ":")

	slport, _ := strconv.Atoi(addresses[1])
	fmt.Println(slport)
	port_chan <- slport
	gs := grpc.NewServer()
	communication.RegisterNodeServer(gs, nodeServer{})
	gs.Serve(listener)
	defer wg.Done()
	defer listener.Close()
}

func main() {
	var wg sync.WaitGroup
	port_chan := make(chan int)
	go Server(&wg, port_chan)
	wg.Add(1)
	slport := <-port_chan
	port := ""
	if len(os.Args) > 1 {
		port = os.Args[1]
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {

			src := &net.TCPAddr{
				IP:   net.ParseIP("0.0.0.0"),
				Port: slport,
			}
			d := net.Dialer{
				Control:   reusePort,
				LocalAddr: src,
			}
			return d.Dial("tcp4", "0.0.0.0:"+port)

		})}
		conn, _ := grpc.Dial("0.0.0.0:"+port, opts...)
		client := communication.NewNodeClient(conn)
		fmt.Println(conn.GetState().String())
		stream, _ := client.InitNode(context.Background())
		stream.Send(&communication.NodeReq{Message: "I am client"})
		fmt.Println(stream.Recv())

	}
	wg.Wait()
}

// static bootstrap node
// node: creates a rsa public private key
// node: here take my public key and store my publicip:port
