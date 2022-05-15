package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net"
	"net/netip"
	"strconv"
	"strings"
	"sync"
	"time"
)

const K_SIZE = 3
const RENDEZVOUS_ADDR = "0.0.0.0:5555"
const BOOTSTRAP_ADDR = "0.0.0.0:6666"
const LOCAL_IP = "0.0.0.0"
const LOCAL_PORT = 6666
const LOCAL_ADDR = LOCAL_IP + ":6666"

var TN *ThisNode

type ResultList struct { //deprecated
	NodeAdd
	Distance float32
	Prev     *ResultList
	Next     *ResultList
}

type KNode struct {
	NodeAdd
	Distance float32
}

type FindNode struct {
	NodeAdd
	KNodes  []KNode
	Queried map[string]bool
}

type ThisNode struct {
	NodeAdd      //type embedding so can call Node.ID()
	Kbuckets     [160][]NodeAdd
	QueriedFlood map[string]bool //for very large net change this
}
type NodeAdd struct {
	IP   string
	PORT int
}

func (node *NodeAdd) DistanceTo(PEERID string) *big.Int {
	thisID := new(big.Int)
	peerID := new(big.Int)
	thisID.SetString(node.ID(), 16)
	peerID.SetString(PEERID, 16)
	distance := new(big.Int)
	distance.Xor(thisID, peerID)
	return distance
}

func (findNode *FindNode) UpdateKNodes(kNodes []KNode) {
	findNodeID := findNode.ID()
	for _, kNode := range kNodes {
		distance := new(big.Float).SetInt(kNode.DistanceTo(findNodeID))
		dvdr := new(big.Float).SetInt(new(big.Int).Lsh(big.NewInt(1), 155))
		kNode.Distance, _ = new(big.Float).Quo(distance, dvdr).Float32()

		big := float32(0.0)
		big_index := 0
		break_stat := false
		for k, node := range findNode.KNodes {
			if node.ID() == kNode.ID() {
				break_stat = true
				break
			}
			if node.Distance > big {
				big = node.Distance
				big_index = k
			}

		}
		if break_stat {
			continue //breaks here but continue from next loop
		}

		if len(findNode.KNodes) < K_SIZE {
			findNode.KNodes = append(findNode.KNodes, kNode)
		} else {
			if kNode.Distance < big || big == 0 {
				findNode.KNodes[big_index] = kNode
			}
		}
	}
	//fmt.Println(findNode.KNodes)
}

func (list *ResultList) Insert(nodes []*ResultList) *ResultList { //deprecated
	for _, node := range nodes {
		for list.Prev != nil {
			list = list.Prev
		}
		for list != nil {
			if node.Distance < list.Distance {
				if list.Prev != nil {
					node.Prev = list.Prev
					list.Prev.Next = node
				}
				list.Prev = node
				node.Next = list

				break

			} else if list.Next == nil {
				list.Next = node
				node.Prev = list
				break
			}
			list = list.Next
		}

	}
	for list.Prev != nil {
		list = list.Prev
	}
	return list

}

func GetPeerIndex(list []NodeAdd, PEERID *string) int {
	for i, nd := range list {
		if nd.ID() == *PEERID {
			return i
		}
	}
	return -1
}

func (node *ThisNode) GetKNodes(PEERID string) []NodeAdd {
	var kNodeList []NodeAdd
	K := K_SIZE
	index := node.GetKbucketIndex(node.DistanceTo(PEERID))
	for i := index; i < 160; i++ {
		if len(node.Kbuckets[i]) == 0 {
			continue
		}
		if len(node.Kbuckets[i]) <= K {
			kNodeList = append(kNodeList, node.Kbuckets[i]...)
			K = K - len(node.Kbuckets[i])

			if K == 0 {
				return kNodeList
			}

		} else {
			fmt.Println("K = ", K)
			kNodeList = append(kNodeList, node.Kbuckets[i][:K]...)
			//K = 0
			return kNodeList
		}

	}
	for i := index - 1; i >= 0; i-- {
		if len(node.Kbuckets[i]) == 0 {
			continue
		}
		if len(node.Kbuckets[i]) <= K {
			kNodeList = append(kNodeList, node.Kbuckets[i]...)
			K -= len(node.Kbuckets[i])
			fmt.Println("K = ", K)
			if K == 0 {
				return kNodeList
			}

		} else {
			fmt.Println("K = ", K)
			kNodeList = append(kNodeList, node.Kbuckets[i][:K]...)
			return kNodeList
		}

	}
	return kNodeList
}
func (node *ThisNode) GetKbucketIndex(distance *big.Int) int {
	distance_float := new(big.Float).SetInt(distance)
	dr128 := new(big.Float).SetInt(new(big.Int).Lsh(big.NewInt(1), 128))
	final_res := new(big.Float).Quo(distance_float, dr128) //result is 32 bits
	num, _ := final_res.Float64()
	return int(math.Log2(num) + 128)
}
func (node *ThisNode) AddToBucket(peer *NodeAdd) {

	K := K_SIZE //max 20 node can be stored in Kbuckets[i]
	PEERID := peer.ID()

	distance := node.DistanceTo(PEERID)
	index := node.GetKbucketIndex(distance)
	//fmt.Println("INDEX = ", index)

	peerIndex := GetPeerIndex(node.Kbuckets[index], &PEERID)
	if peerIndex == -1 && len(node.Kbuckets[index]) < K { //peer is not there and size is < K
		node.Kbuckets[index] = append(node.Kbuckets[index], *peer)
	} else if peerIndex != -1 && len(node.Kbuckets[index]) <= K { //peer is there but size is <= K
		node.Kbuckets[index] = append(node.Kbuckets[index], *peer)
		node.Kbuckets[index] = append(node.Kbuckets[index][:peerIndex], node.Kbuckets[index][peerIndex+1:]...)
	} else if peerIndex == -1 && len(node.Kbuckets[index]) == K { //peer is not there but size if full
		//ping the first node if replies ok then move it to the last and reject this peer
		//is ping reponse is failed then delete the first node and then append the peer
	}

}
func (node *NodeAdd) ID() string {
	address := node.IP + fmt.Sprint(node.PORT)
	h := sha1.New()
	h.Write([]byte(address))
	sm := h.Sum(nil)
	sms := fmt.Sprintf("%x", sm)
	return sms
}
func RunUDPServer(wg *sync.WaitGroup) {
	conn, _ := net.ListenPacket("udp4", LOCAL_ADDR) //conn is done to reserve a socket for this, I can write to different address using this same socket
	addresses := strings.Split(conn.LocalAddr().String(), ":")
	slport, _ := strconv.Atoi(addresses[1])
	fmt.Println(slport)
	var rwg sync.WaitGroup
	ping_quit := make(chan string)
	go Read(&rwg, conn, ping_quit)
	rwg.Add(1)

	//Making a ThisNode
	TN = &ThisNode{NodeAdd: NodeAdd{IP: LOCAL_IP, PORT: LOCAL_PORT}, QueriedFlood: map[string]bool{}}

	ticker := time.NewTicker(5 * time.Second)
	go PingToRendezvous(conn, ticker, ping_quit) //every n time

	rwg.Wait()
	defer wg.Done()
	defer conn.Close()
}
func (node *NodeAdd) LoadFromString(addr string) {
	addresses := strings.Split(addr, ":")
	slport, _ := strconv.Atoi(addresses[1])
	node.IP = addresses[0]
	node.PORT = slport
}
func Read(wg *sync.WaitGroup, conn net.PacketConn, ping_quit chan string) {
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
		fmt.Println(out["RPC"])
		if out["RPC"] == "PING_OK" {
			ping_quit <- "ok"
		}
		if out["RPC"] == "FIND_NODE" {
			kNodes := TN.GetKNodes(out["NODE_ID"].(string)) //type assertion
			nd := &NodeAdd{}
			nd.LoadFromString(addr.String())
			TN.AddToBucket(nd)
			fmt.Println(kNodes)
			payload, _ := json.Marshal(map[string]any{"RPC": "FIND_NODE_RES", "NODE_ID": out["NODE_ID"], "KNODES": kNodes})
			Write(conn, addr.String(), payload)
			//fmt.Println(TN.Kbuckets)
		}
		if out["RPC"] == "FLOOD" {
			TN.SendFloodTaskToAllExcept(conn, addr.String(), out["TASK"].(string), out["VALUE"].(string), out["SOURCE_ADDR"].(string))
		}
		if out["RPC"] == "CONNECT_DO" {
			payload, _ := json.Marshal(map[string]string{"RPC": "ALIVE"})
			Write(conn, fmt.Sprint(out["NODE_ADDR"]), payload)
		}
	}

}
func Write(conn net.PacketConn, dst string, payload []byte) {
	//defer conn.Close()
	//someone should call it
	addr, _ := netip.ParseAddrPort(dst)
	dest_addr := net.UDPAddrFromAddrPort(addr)
	conn.WriteTo(payload, dest_addr)

}

func PingToRendezvous(conn net.PacketConn, ticker *time.Ticker, quit chan string) {
	//defer close(quit)
	for {
		select {
		case <-ticker.C:
			for i := 0; i < 1; i++ {
				payload, _ := json.Marshal(map[string]any{"RPC": "PING"})
				Write(conn, RENDEZVOUS_ADDR, payload)

			}

		case <-quit:
			//ticker.Stop()
			break
		}
	}
}

func main() {

	/*
		lport := 54635
		thisNode := ThisNode{NodeAdd:NodeAdd{IP:"127.0.0.1",PORT:lport}}
		peer1:=NodeAdd{IP:"127.0.0.1",PORT:51835}
		thisNode.AddToBucket(&peer1)


		fNode := FindNode{NodeAdd:NodeAdd{IP:"3232",PORT:432423}}
		fNode.UpdateKNodes([]KNode{
			KNode{NodeAdd:NodeAdd{IP:"124",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"43253",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"153524",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:6342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"57438990",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:442}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},
			KNode{NodeAdd:NodeAdd{IP:"15454324",PORT:4342}},

		})*/
	var wg sync.WaitGroup
	go RunUDPServer(&wg)
	wg.Add(1)
	wg.Wait()

}

func (thisNode *ThisNode) SendFloodTaskToAllExcept(conn net.PacketConn, exceptAddr string, task string, value string, sourceAddr string) {
	payload, _ := json.Marshal(map[string]string{"RPC": "FLOOD", "SOURCE_ADDR": sourceAddr, "TASK": task, "VALUE": value})
	sr_nd := &NodeAdd{}
	sr_nd.LoadFromString(sourceAddr)
	for _, list := range thisNode.Kbuckets {
		for _, dst := range list {
			if dst.ID() != exceptAddr && dst.ID() != sr_nd.ID() && !TN.QueriedFlood[dst.ID()] {
				Write(conn, dst.IP+":"+fmt.Sprint(dst.PORT), payload)
				TN.QueriedFlood[dst.ID()] = true
			}
		}
	}
}
