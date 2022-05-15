package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const K_SIZE = 3

const RENDEZVOUS_ADDR = "13.233.39.180:5555"
const BOOTSTRAP_ADDR = "13.233.39.180:6666"
const LOCAL_IP = "0.0.0.0"

var FN *FindNode
var TN *ThisNode

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
func (node *NodeAdd) LoadFromString(addr string) {
	addresses := strings.Split(addr, ":")
	slport, _ := strconv.Atoi(addresses[1])
	node.IP = addresses[0]
	node.PORT = slport
}
func (node *NodeAdd) ID() string {
	address := node.IP + fmt.Sprint(node.PORT)
	h := sha1.New()
	h.Write([]byte(address))
	sm := h.Sum(nil)
	sms := fmt.Sprintf("%x", sm)
	return sms
}

type ThisNode struct {
	NodeAdd      //type embedding so can call Node.ID()
	Kbuckets     [160][]NodeAdd
	QueriedFlood map[string]bool //for very large net change this
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
	distance_float := new(big.Float).SetInt(distance)
	dr128 := new(big.Float).SetInt(new(big.Int).Lsh(big.NewInt(1), 128))

	final_res := new(big.Float).Quo(distance_float, dr128) //result is 32 bits
	num, _ := final_res.Float64()
	index := int(math.Log2(num) + 128)

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

type KNode struct {
	NodeAdd
	Distance float32
}

type FindNode struct {
	NodeAdd
	KNodes  []KNode
	Queried map[string]bool
}

func (findNode *FindNode) SendFindNodeToAllKNodes(conn net.PacketConn) {
	payload, _ := json.Marshal(map[string]any{"RPC": "FIND_NODE", "NODE_ID": findNode.ID()}) //FIND_NODE_RES
	for _, dst := range findNode.KNodes {
		//fmt.Println(dst)
		if !FN.Queried[dst.ID()] {
			Write(conn, dst.IP+":"+fmt.Sprint(dst.PORT), payload)
		}

	}

}
func (findNode *FindNode) UpdateKNodes(kNodes []KNode) bool {
	changed := false
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
			changed = true
		} else {
			if kNode.Distance < big || big == 0 {
				findNode.KNodes[big_index] = kNode
				changed = true
			}
		}
	}
	//fmt.Println(findNode.KNodes)
	return changed
}

func GetPeerIndex(list []NodeAdd, PEERID *string) int {
	for i, nd := range list {
		if nd.ID() == *PEERID {
			return i
		}
	}
	return -1
}

func Write(conn net.PacketConn, dst string, payload []byte) {
	//defer conn.Close()
	//someone should call it

	// send 2 CONNECT request

	addr, _ := netip.ParseAddrPort(dst)
	dest_addr := net.UDPAddrFromAddrPort(addr)
	//fmt.Println("DST = ", dest_addr)
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

func Read(wg *sync.WaitGroup, conn net.PacketConn, ping_quit chan string, stunAddr chan string) {
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

		if out["RPC"] == "PING_OK" {
			ping_quit <- "ok"
		}
		if out["RPC"] == "FIND_NODE" {
			kNodes := TN.GetKNodes(out["NODE_ID"].(string)) //type assertion
			nd := &NodeAdd{}
			nd.LoadFromString(addr.String())
			TN.AddToBucket(nd) //add to kbucket
			//fmt.Println(kNodes)
			payload, _ := json.Marshal(map[string]any{"RPC": "FIND_NODE_RES", "NODE_ID": out["NODE_ID"], "KNODES": kNodes})
			Write(conn, addr.String(), payload)
			//fmt.Println("K-BUCKETS =", TN.Kbuckets)
		}
		if out["RPC"] == "FIND_NODE_RES" {
			fmt.Println("FIND NODE RES")
			//fmt.Println(addr.String())
			if FN.ID() == out["NODE_ID"] {
				sender := &NodeAdd{}
				sender.LoadFromString(addr.String())
				FN.Queried[sender.ID()] = true //it helps to keep track of previously requested peer for current FN
				if out["KNODES"] != nil {
					_kNodes := out["KNODES"].([]interface{})
					var kNodes []KNode
					for _, obj := range _kNodes {
						kn := &KNode{}
						new_obj := obj.(map[string]any)
						ip := fmt.Sprint(new_obj["IP"])
						port := fmt.Sprint(new_obj["PORT"])
						kn.LoadFromString(ip + ":" + port)
						if kn.ID() != TN.ID() { //should not appeneded
							TN.AddToBucket(&kn.NodeAdd)
							kNodes = append(kNodes, *kn)
						}
					}
					changed := FN.UpdateKNodes(kNodes)
					if changed {
						FN.SendFindNodeToAllKNodes(conn)
						changed = FN.UpdateKNodes(kNodes)
					}
				}
			}

			//fmt.Println(FN.KNodes)

		}
		if out["RPC"] == "FLOOD_RES" {
			fmt.Println("RESULT IS  =", out["RESULT"])
		}

		if out["RPC"] == "FLOOD" {
			TN.SendFloodTaskToAllExcept(conn, addr.String(), fmt.Sprint(out["TASK"]), fmt.Sprint(out["VALUE"]), fmt.Sprint(out["SOURCE_ADDR"]))
			var val int
			fmt.Sscanf(fmt.Sprint(out["VALUE"]), "%d", &val)
			result := fmt.Sprint(val / 2)
			payload, _ := json.Marshal(map[string]string{"RPC": "FLOOD_RES", "RESULT": result})
			fmt.Println("I CALCULATED SOMETHING!")
			Write(conn, out["SOURCE_ADDR"].(string), payload)

		}
		if out["RPC"] == "ALIVE" {
			//fmt.Println(out["RPC"], " ", addr.String())
		}
		if out["RPC"] == "CONNECT_DO" {
			payload, _ := json.Marshal(map[string]string{"RPC": "ALIVE"})
			Write(conn, fmt.Sprint(out["NODE_ADDR"]), payload)
			//Write(conn, fmt.Sprint(out["NODE_ADDR"]), payload)
			//Write(conn, fmt.Sprint(out["NODE_ADDR"]), payload)
			//fmt.Println(out["RPC"], " ", addr.String())
		}
		if out["RPC"] == "STUN_RES" {
			_stunAddr := fmt.Sprint(out["NODE_ADDR"])
			//fmt.Println("STUN ADDR = ", _stunAddr)
			stunAddr <- _stunAddr
		}
	}

}
func RunUDPServer(wg *sync.WaitGroup, connection chan net.PacketConn, stunAddr chan string) {

	conn, _ := net.ListenPacket("udp4", LOCAL_IP+":") //conn is done to reserve a socket for this, I can write to different address using this same socket
	addresses := strings.Split(conn.LocalAddr().String(), ":")
	slport, _ := strconv.Atoi(addresses[1])
	fmt.Println(slport)
	connection <- conn

	//Making a ThisNode
	TN = &ThisNode{NodeAdd: NodeAdd{IP: LOCAL_IP, PORT: slport}, QueriedFlood: map[string]bool{}}

	//making a FindNode
	FN = &FindNode{NodeAdd: TN.NodeAdd, Queried: map[string]bool{}}

	//adding bootstrap node to FindNode's latest K-NODES
	bn := &NodeAdd{}
	bn.LoadFromString(BOOTSTRAP_ADDR)
	TN.AddToBucket(bn)

	kn := &KNode{}
	kn.LoadFromString(BOOTSTRAP_ADDR)
	FN.UpdateKNodes([]KNode{*kn})

	//adding bootstrap node to the TN's KBucket is not needed since added to FindNode

	//sending FIND_NODE RPC to All FindNode's K-NODES
	FN.SendFindNodeToAllKNodes(conn)

	var rwg sync.WaitGroup
	ping_quit := make(chan string)

	go Read(&rwg, conn, ping_quit, stunAddr)
	rwg.Add(1)

	rendezvousTicker := time.NewTicker(5 * time.Second)

	go PingToRendezvous(conn, rendezvousTicker, ping_quit) //every n time

	// ping to all bucket Nodes
	bucketTicker := time.NewTicker(5 * time.Second)
	bucketTickQuit := make(chan string)
	go TN.PingToBuckets(conn, bucketTicker, bucketTickQuit)

	rwg.Wait()
	defer wg.Done()
	defer conn.Close()
}
func main() {
	var wg sync.WaitGroup
	connection := make(chan net.PacketConn)
	stunAddr := make(chan string)
	go RunUDPServer(&wg, connection, stunAddr)
	wg.Add(1)
	conn := <-connection

	payload, _ := json.Marshal(map[string]string{"RPC": "STUN"})
	Write(conn, RENDEZVOUS_ADDR, payload)
	sAddr := <-stunAddr

	if len(os.Args) > 1 {
		prime_num := os.Args[1]
		exceptAddr := "" //the addrID from where the FLOOD_RES received
		//source := TN.NodeAdd //who initiated the message first
		task := "PRIME" //a prime number
		value := prime_num

		//TN.SendFloodTaskToAllExcept(conn, exceptAddr, task, value, source.IP+":"+fmt.Sprint(source.PORT))
		//time.Sleep(40 * time.Second)
		TN.SendFloodTaskToAllExcept(conn, exceptAddr, task, value, sAddr)

	}
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

func (thisNode *ThisNode) PingToBuckets(conn net.PacketConn, ticker *time.Ticker, quit chan string) {
	//defer close(quit)
	for {
		select {
		case <-ticker.C:
			for _, list := range thisNode.Kbuckets {
				for _, dst := range list {
					payload, _ := json.Marshal(map[string]string{"RPC": "ALIVE"})
					Write(conn, dst.IP+":"+fmt.Sprint(dst.PORT), payload)
					payload, _ = json.Marshal(map[string]string{"RPC": "CONNECT", "NODE_ADDR": dst.IP + ":" + fmt.Sprint(dst.PORT)})
					Write(conn, RENDEZVOUS_ADDR, payload)

				}
			}

		case <-quit:
			//ticker.Stop()
			break
		}
	}
}

/*
first ask bootstrap node to send the best KNodes closest to a node of itself.

*/

//same ip nat traversal maynot work, try different internet
