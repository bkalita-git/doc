package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net"
	"net/netip"
	"strconv"
	"strings"
	"sync"
)

const LOCAL_IP = "0.0.0.0"
const LOCAL_PORT = 5555
const LOCAL_ADDR = LOCAL_IP + ":5555"

type NodeAdd struct {
	IP   string
	PORT int
}

func (node *NodeAdd) ID() string {
	address := node.IP + fmt.Sprint(node.PORT)
	h := sha1.New()
	h.Write([]byte(address))
	sm := h.Sum(nil)
	sms := fmt.Sprintf("%x", sm)
	return sms
}

func Read(wg *sync.WaitGroup, conn net.PacketConn) {
	Data := make(map[string]NodeAdd)
	defer conn.Close()
	defer wg.Done()
	//continuous
	for {
		payload := make([]byte, 1024)
		n, addr, _ := conn.ReadFrom(payload)
		var out map[string]any
		err := json.Unmarshal([]byte(string(payload[:n])), &out)
		if err != nil {
			fmt.Println(err)
		}

		if out["RPC"] == "PING" {
			payload, _ := json.Marshal(map[string]any{"RPC": "PING_OK"})
			Write(conn, addr.String(), payload)
			addresses := strings.Split(addr.String(), ":")
			slport, _ := strconv.Atoi(addresses[1])
			peerNode := &NodeAdd{IP: addresses[0], PORT: slport}
			Data[peerNode.ID()] = *peerNode
			//fmt.Println(Data)

		}

		if out["RPC"] == "CONNECT" {
			payload, _ := json.Marshal(map[string]string{"RPC": "CONNECT_DO", "NODE_ADDR": addr.String()})
			Write(conn, fmt.Sprint(out["NODE_ADDR"]), payload)
			fmt.Println(out["RPC"], " ", addr.String(), " SENT TO ", fmt.Sprint(out["NODE_ADDR"]))
		}

		if out["RPC"] == "STUN" {
			payload, _ := json.Marshal(map[string]string{"RPC": "STUN_RES", "NODE_ADDR": addr.String()})
			Write(conn, addr.String(), payload)
		}
		//fmt.Println(addr.String(), "  ", string(payload[:n]), "  ", out)
	}

}
func Write(conn net.PacketConn, dst string, payload []byte) {
	//defer conn.Close()
	//someone should call it
	addr, _ := netip.ParseAddrPort(dst)
	dest_addr := net.UDPAddrFromAddrPort(addr)
	conn.WriteTo(payload, dest_addr)

}

func main() {
	conn, _ := net.ListenPacket("udp4", LOCAL_ADDR) //conn is done to reserve a socket for this, I can write to different address using this same socket
	addresses := strings.Split(conn.LocalAddr().String(), ":")
	slport, _ := strconv.Atoi(addresses[1])
	fmt.Println(slport)
	var wg sync.WaitGroup
	go Read(&wg, conn)
	wg.Add(1)
	wg.Wait()
	defer conn.Close()
}
