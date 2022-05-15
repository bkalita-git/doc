Kademlia is Distributed Hash Table for Decentralized Network.
contacts only O(log(n)) nodes out of n nodes.

First Generation P2P Networks such as Napster relied on a central database to co-ordinate look ups on the network.
Second Generation P2P Netwotks such as Gnutella used flooding.
3rd gen such as Bittorrent use DHT

NODE ID can be UUID(random) but still geographically can be similar for a country's nodes that other country's node.
XOR can be used because it acts as a distance function between node IDs
and distance between node itself is zero. also it is symmetric, A TO B and B TO A are same. also it follows traingle ineqality, i.e. distance from A to B is shorter than the sum of B to C and C to B

A basic Kademlia network with 2n nodes will only take n steps (in the worst case) to find that node.

if a node id is 128 bits then that node will keep 128 list where 128 nodes info will be stored.
an index will have
{
node server ip [grpc ip]
node server port[grpc port]
node UUID[128 bits]
routine table
api to call other nodes
}

routing table{
a list of buckets having 128 length[store info of other nodes]
}

-------------------128 bits-----------------------
first node id[128-1 bits]

let's say n=3 so UUID can have 8 node in the network.
first let's say node 0 is global ip[bootstrap]
node1[give me my closest node]->node0[Oh, I don't have any , you know what, I will keep your info to my routing table]->node1[hmm, node0 responds me so now i have to store node0]
node2[give me my closest node]->node0[i will keep you and take node1]->[node2,hmm i will keep node0,node1, hey node1 are you alive? ping]->node1[yes i am, oh wait I AM adding you too.]

key:value[128 bit]

The distance mentioned here is a logical distance, which has nothing to do with the geographical location. Therefore, it is possible that the calculated logical distance between two nodes is very close, but in fact, the geographical distance is very far.

n subtree for each n node.
from each subtree k nodes will be selected.

for a node, n routine tables will be mainrained.
1 subtree == 1 routing table == 1 list == 1 k-bucket == containts K items
1 node == n subtrees == n routine table == n list == n k-bucket

MOST BASIC:
DISTRIBUTED HASH TABLE:
("key","value") pair ==> hash("key")=hashvalue=array_index="value"
what if we do this- hash("value") = "key" and then send ("key","value") in a distributed network and store only by the specific node. the thing is what specific node? there must be another mapping from "key"->node and this is achieved by constistent hashing or rendezvous hasing.
so, we have also NODE ID.
if the distance of NODE ID "98r43feij" and the "key" is minimum then that ("key","value") should be stored by that node.

## Algorithm

We have k=20, alpha=2, hash_size=160

- every node generates a random 160 GUID(global user Id)
- every node has 160 k-buckets/lists. they are indexed from k-bucket[0] to k-bucket[159] and each k-bucket is one list.
- for Node ID 0(converted 160 bits to decimal) its k-bucket[i] list will hold Other Node ID&Info where Distance(0,Other_node_id) falls in between $[2^i,2^{i+1})$. but each list can contain maximum k node information. so `max(len(k-bucket[i]))=k`
  ```
  [1,2,3,4) = [1,2,3]
  ```
- The main Aim is that `Node ID A` tries to find/gather k(a list of size k,here 20) closests node to a `Node ID B` and send search request to all k but alpha nodes at a time.
- The aplha(systemwide concurrency parameter) is the selected number of nodes to which find_node_B will be sent parallely,
- response from the k nodes will be added again to the list because we need the best 20 nodes. and repeat this process till no changes in the list.

## REFERENCES

1. https://en.wikipedia.org/wiki/Distributed_hash_table [must read first]
1. https://developpaper.com/kademlia-protocol/
1. https://sci-hub.se/10.1109/3PGCIC.2013.15
