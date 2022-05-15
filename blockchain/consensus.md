## Consensus - A distributed Agreement Mechanism

Let's say there were no consensus. Now in a distributed network there are 10 nodes. Now say a user (outside the 10 nodes) sends a data(transaction) to any one node say NODE_B. Now NODE_B will distribute that data so that the result is All the 10 nodes have it. Now every nodes want to add the data as a block to the blockchain, and lets say each one added it. but when any one node appends a block to the blockchain it then tells other nodes to add that block. but here we got two problems -

1.  since every node will add the block so every node tells others to add the block but we only need 1 node to tell the other nodes.
2.  even if NODE_B tells other 9 nodes to add the block to their own blockchain, Those 9 nodes can't just add it because there is no distributed rule followed by NODE_B while adding it and there's no proof of that rule.

So, Basically we want a distributed rule and a proof so others can verify that the proof tells that "the node followed the rule while adding the block."

such rules can be -

1. "Every node should compute a difficulty hash" - proof of work
2. "Every node agrees to a same decision that they allows the initiator to add the block or reject the block" - slush Consensus

But why do we need the distributed rule or the distributed agreement?
First of all no Node can alter user data, since the user will sign his/her data with the private key. so Let's say NODE_B is malicious node and alters the user data and added to the blockchain and tell other 9 nodes to do so. But other 9 nodes can verify the data using the user's public key so it could detect if correct one or not. So, basically **a consensus or a distributed rule is there to verify that the node who tells other nodes to add a block is the node of same type, that is all nodes haves the same software and each verifies the transaction in the same way**.

## But poor consensus is bad

For example, if the difficulty level of hash in proof of work is very low then some one could alter the blockchain and if this is done to 51% machine then the blockchain will have false data.

## Consensus is different from transaction verification

consensus is needed to add the block/data to the blockchain.

## Avalanche Consensus(frosty at higher level)

Level of consensus in Avalanche.

```
slush --> snowflake --> snowball --> avalanche --> snowman --> frosty
```

the core of the consensus in avalanche snowman

- Slush consensus  
  Trying the query accept/reject value to k random nodes for by sending the transaction.
- Snowflake consensus  
  To prevent malicious node's decision for a transaction by keeping consecutive same results.
- Snowball consensus
  Keeps track of confidence level of 1 transaction if the node should accept/reject the transaction since mallicious node can flip its preference.

* Avalanche consensus
  It maintains a Dag instead of chain. Batches multiple transactions into a vertex and keep tracks the confidence. Here vertex is an instance of snowball consensus that means 1 snowball consensus has many transaction. It's to prevent conflict transaction. uses **chit** values
* Snowman  
  Instead of DAG it's a chain.
* Frosty  
  Coming soon...

Apart from that a node can only participate in the consensus if it stakes 2000AVAX, this prevents malicious node from participating the network. And this is done by API call by a user. So if a user write a transaction like "I want to be a validator here my 2000AVAX staking to this Node ID" and then that user became a validator and now that node-id can join the consensus protocol.

## Double Spending

sending the same money twice. for example I have 1$ and I made two transaction to send 1$ to bob and alice , but I have only 1$.
Suppose two different miners will pick both transactions at the same time(ie. timestamp of the two block are same somehow) and start creating a block. Now, when the block is confirmed, both Alice and Bob will wait for confirmation on their transaction. Whichever transaction first got confirmations will be validated first, and another transaction will be pulled out from the network.
Now suppose if both Alice and Bob received the first confirmation at the same time, then there is a race will be started between Alice and Bob. So, whichever transaction gets the maximum number of confirmations from the network will be included in the blockchain, and the other one will be discarded.

## Avalanche DAG

- Verification of a transaction
- Different applications will have different notions about what it means for two transactions to conflict.

chit : if a node query to k nodes "if tx should be prefered?" and if alpha number out of k nodes response "yes" then the transaction "tx" get a "chit=1"  
confidence(tx) = chit(tx) + descendent_txs_chit(tx)
consecutive_success(tx) = consecutive_success(tx) + descendent_txs_consecutive_success(tx)
consecutive successses = No of continous "yes" response from alpha nodes.

Ultimately, if 1 transaction gets a chit then the node has to update its ancestors confidence and consecutive successes. for a transaction tx its immediate ancestor's consecutive success is set to 0 iff tx does not get a chit.
Now, In a Dag(say 10 transactions are there)whichever "tx" gets a consecutive success >= beta, that tx is accepted by the system.
Also, A Node having a DAG and in it a conflict set where 3 transactions are there, The transaction having the highest confidence will be preferred by the Node.
