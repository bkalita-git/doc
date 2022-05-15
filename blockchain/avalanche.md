- It's all about Nodes and they communicate with each other.
- A block is an instance of a VM, and chain is where previous blocks are stored.
- from the same VM multiple blockchain can be made.
- A VM can be run as a sperate process from AvalancheGo and so a running validator can run a custom VM using its compiled binary.
- Every node is in a subnet. The primary network is also a subnet.
- A node can perticipate in different subnet.
- 1 blockchain can be validated by using only 1 subnet. but 1 subnet can validate many blockchain.
- A VM is a custom server providing services(through REST API) you define. for example p-chain,c-chain,x-chain are also 3 VM.
- A Node, VM, Subnet, Blockchain, Block, Asset, Transaction, Transaction_type has ID.
- P and C chain uses Snowman consensus and X chain uses Avalanche consensus. Only validators(who stakes a lot) can participate in the consensus, other nodes can have the blockchain and read/write it but can't validate.
- 1 NODE MULTIPLE VM, 1 VM 1 CONSENSUS MULTIPLE BLOCKCHAIN/DAG
- Every transaction shouldbe generated in offline(javascript) mode and only use the NODE_API_CALL to send the transaction.
- To create a new Blockchain (ie instance of a VM) you must provide the Genesis data for the blockchain.
- genesis data is built using a particular(can be a custom VM) VM api but the creation of the custom blockchain is done using platform VM(by providing subnet id and VM id and that genesis data) but adding data(block) to the blockchain or exploring it is done through the Custom VM api.
- bitcoin and avalanche and some multichain uses UTXO MODEL

## Node Implementation of Avalanche

- https://github.com/ava-labs/avalanchego

## We need a cluster of `avalanchego` node without a centralized server

1. TestNet (use fuji load balancer)
2. MainNet (use mainnet load balancer)
3. Local (using avalanche-network-runner to build the cluster and it will provide endpoints to every node)

   ```
   avalanche-network-runner server --log-level debug --port=":8081" --grpc-gateway-port=":8082"

   ```

   ```
   avalanche-network-runner control start --log-level debug --endpoint="0.0.0.0:8081" --avalanchego-path /home/common/Programs/Arch_linux/avalanchego/build/avalanchego
   ```

## Public Api

Multiple AvalancheGo(server) Nodes running and connected to a LoadBalancer and is publicly available.

- Avalanche Mainnet Endpoint

  - https://api.avax.network/:443
  - For C-Chain API, the URL is https://api.avax.network/ext/bc/C/rpc.(EVM)
  - For X-Chain API, the URL is https://api.avax.network/ext/bc/X. (AVM,wallet creation)
  - For P-Chain API, the URL is https://api.avax.network/ext/P (Platform VM, provides creation on subnet)

- Avalanche Testnet Endpoint
  - https://api.avax-test.network:443
- How to use [here](https://docs.avax.network/build/avalanchego-apis/issuing-api-calls/)

## Virtual Machine

like in bitcoin those software are static, follows one rule to validate one blockchain. So, multiple computers installed the same software and then they followed the same rules like calculating the valid hash of the transaction. For example, I implemented a distributed network using Kademlia algorithm. It's like the first step of blockchain network.

Now, what if in the blockchain we don't want to store the transaction, so we want custom blockchain for which we want to validate/add so the software should be different now. That's where we need a solution where we can define custom Blockchain and then we can define a VM(set of rules for the blockchain) for that custom blockchain and that VM will get executed by the software. And that software is `AvalancheGo`

creation of custom blockchain is same as first creating the custom VM. and a custom BlockChain is just an instance of that custom VM. we can write our custom VM in Golang. [here](https://docs.avax.network/build/tutorials/platform/subnets/create-a-virtual-machine-vm)

A virtual machine defines Blockchain's state, state transition function, transaction, and API s through which users can interact. We don't need to care about the underlying network setup since `AvalancheGo` will take care of that.

`rpcchainvm` is a special type of VM that wraps user defined VM. `rpcchainvm` has a server part in which the user defined VM runs as a process separate from `AvalancheGo` process and allows the underlying VM's methods to be called via gRPC. and `rpcchainvm` client parts runs with the `AvalancheGo` and makes gRPC calls to the server part.

```
AvalancheGo calls rpcchainvm_client_method <--gRPC--> rpcchainvm calls user_defined_VM's_Method
```

That means, `EVM` can also be wrapped inside `rpcchainvm` with slight modification and so `AvalancheGo` can make gRPC calls.

`AVM` stands for Avalanche Virtual Machine which is a built in VM and does same by the `rpcchainvm`

for linear blockchain (VM) avalanchego uses its snowman consensus engine. There can be DAG Blockchain(VM) also.

By default `AvalancheGo` node comes with 3 blockchains, p-chain, x-chain, c-chain. So, there must be VM associated with them. x-chain uses `AVM`, c-chain uses `EVM`.

## Validators and delegators (these are user who has addresses and issue some special type of transaction on P chain)

Now, what is a Primary Network? It's a network of validators(not delegators). and each Validator is just a node/ec2 runnning `AvalancheGo` which deposited 2000AVAX to the network. Now, these validators maintain the Primary Network means these validators validate all the 3 blockchain x-chain, p-chain, c-chain. Now someone can call REST api to any one validator to create a wallet. so avalanche does this by putting a loadbalancer and then it selects a validator to send the request.

delegators on the other hands is a user who stake(bondhok t tho) some AVAX to a validator.

## Private Blockchain network/ Public Blockchain network

for example, in Avalanche Mainnet my laptop can join by running `AvalancheGo` and get the blockchain(130 GB say) and start validating new blocks and starts adding. SO, basically in Public blockchain network anyone can join the network as validator/miner nodes.

But in Private Blockchain Network the network has pre defiened information for who can become validator/miner.

so, For example, we created a custom blockchain(so custom VM) and so only those nodes can process our blockchain who has our custom VM installed.

## Creating a Private Subnet from the Mainnet

[creating subnet](https://docs.avax.network/build/tutorials/platform/subnets/create-a-subnet/)

## Running a Node

```
git clone https://github.com/ava-labs/avalanchego.git
cd avalanchego
./scripts/build.sh #this will create executable named "avalanchego"
./avalanchego #this runs the Avalanche Node. try --help for more options
```

- running by Connecting to `MainNet`

  ```
  ./avalanchego #this will try to connect with the MainNet and tries to download blocks from the p-chain.
  ```

- running without connecting to any Net
  ```
  ./avalanchego --network-id=local --public-ip=127.0.0.1 --http-port=9650 --staking-port=9651 --db-dir=db/node1
  ```
- running by connecting to a test net (fuji for example). This will bootstrap a fewer blocks.

  ```
  ./avalanchego --network-id=fuji
  ```

  So whenever we try to connect to a Network, whether mainnet or testnet then it tries to pull the 3 blockchains(x,p,c) hence does the bootstraping process.

  When we run a node then we get -

  - NodeID-U9CLKTqSDyja5KMqhwA3xd75kx61uUMb
  - routes -
    - /ext/metrics
    - /ext/health
    - /ext/health/readiness
    - /ext/health/health
    - /ext/health/liveness
    - /ext/vm/rWhpuQPF1kb72esV2momhMuTYGkEb1oL29pt2EBXWmSy4kxnT #platform VM
    - /ext/vm/jvYyfQTxGMJLuGWa55kdP2p2zSUYsQ5Raupu4TW34ZAUBAbtq #AVM
    - /ext/vm/mgj786NP7uDwBCcq6YwThhaN8FLyybkCa4zBWTQbNgmK6k9A6/rpc #EVM
    - /ext/info
    - /ext/bc/2CA6j5zYzasynPsFeNoqWkmTCt3VScMvXUZHbfDJ8k3oGzAPtU/avax
    - /ext/bc/2CA6j5zYzasynPsFeNoqWkmTCt3VScMvXUZHbfDJ8k3oGzAPtU/rpc
    - /ext/bc/2CA6j5zYzasynPsFeNoqWkmTCt3VScMvXUZHbfDJ8k3oGzAPtU/ws
    - /ext/bc/11111111111111111111111111111111LpoYY
    - /ext/bc/2eNy1mUFdmaxXNj1eQHUe7Np4gju9sJsEtWQ4MX3ToiNKuADed/wallet
    - /ext/bc/2eNy1mUFdmaxXNj1eQHUe7Np4gju9sJsEtWQ4MX3ToiNKuADed/events
    - /ext/bc/2eNy1mUFdmaxXNj1eQHUe7Np4gju9sJsEtWQ4MX3ToiNKuADed

- creates 3 chains -

  - ```
    #P-CHAIN
    ID: 11111111111111111111111111111111LpoYY
    VMID:rWhpuQPF1kb72esV2momhMuTYGkEb1oL29pt2EBXWmSy4kxnT
    ```

  - ```
    #C-CHAIN
    ID: 2CA6j5zYzasynPsFeNoqWkmTCt3VScMvXUZHbfDJ8k3oGzAPtU
    VMID:mgj786NP7uDwBCcq6YwThhaN8FLyybkCa4zBWTQbNgmK6k9A6
    ```
  - ```
    #X-CHAIN
    ID: 2eNy1mUFdmaxXNj1eQHUe7Np4gju9sJsEtWQ4MX3ToiNKuADed
    VMID:jvYyfQTxGMJLuGWa55kdP2p2zSUYsQ5Raupu4TW34ZAUBAbtq
    ```

* Running another node which bootstrap from another running Node to build a network.
  - ```
    ./build/avalanchego --public-ip=127.0.0.1 --http-port=9650 --staking-port=9651 --db-dir=db/node1 --network-id=local --staking-tls-cert-file=$(pwd)/staking/local/staker1.crt --staking-tls-key-file=$(pwd)/staking/local/staker1.key
    ./build/avalanchego --public-ip=127.0.0.1 --http-port=9652 --staking-port=9653 --db-dir=db/node2 --network-id=local --bootstrap-ips=127.0.0.1:9651 --bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg --staking-tls-cert-file=$(pwd)/staking/local/staker2.crt --staking-tls-key-file=$(pwd)/staking/local/staker2.key
    ```
  * Same thing can be done using "NetworkRunner"
* Create a subnet https://docs.avax.network/build/tutorials/platform/subnets/create-a-subnet/
* More on https://docs.avax.network/build/tutorials/platform/create-a-local-test-network/
* Now use Postman Collections provided by the Avalanche Team.

## Address

An Asymmetric key-pair (Private + Public) is required. To make that someone can use RSA or others algorithm. Most blockchain platform does this by Elliptic curve multiplication(ECM). So, for that A random number is generated which will be the Private Key. Now the Private key is fed to the ECM Algo which generates Public Key. Now that Public Key is hashed to form an address.
What we can do is we can generate the Random number (Private Key) using True Number Generator.

## Smart Contract

Like transactions , Smart Contract is consists of Compiled Code with its hash value as unique ID and stored on the blockchain. So others can invoke it by its ID. So the address of the smart contract is its unique hash value.

## Avalanche Transactions vs Blocks

A block can contain thousand of transactions. But in Avalanche 1 block is 1 transaction. The X blockchain in Avalanche is not a linear list of blocks, It's a DAG of transaction called Vertex. But P chain is a conevntional blockchain where transactions are inside block. P chain is just a linear version of X chain.

## Assets

Bitcoin, Avax, Ether are just assets created on a blockchain. The Avax asset is created on the X-Chain which is instance of Avalanche Virtual Machine. Also, there are many assets created on X-Chains.

## User

So, your running AvalancheGo node is just a server and you want to provide services to users. SO, someone can create a user on your node by calling createuser api. your node also support auth mechanism so you can restrict any REST API by using token. Now, that user can create addresses on blockchain avalilable on your node. for example that user can create X-chain, P-chain, C-Chain `address` on your node. User can export a private key againsts the `address` of a particular chain. Now, another different user can gain access/control of that address by importing the privatekey.So, Also the address has the publickey encapsulated in it. An address can hold a balance of cryptocurrencies. A user can create his own blockchain using AVM using "platform.createBlockchain" api call with json data including the P-Chain address for the user on the P-CHAIN VM service.

on X-Chain user can also create assets. assets have id.

## Wallet

An address can hold a balance of cryptocurrencies. A wallet controls a set of addresses. Think of an address like a lockbox, and a wallet as a key for many lockboxes.

By default creation of 1 wallet will create 1 user/account which is called Main account. and We can add more accounts to 1 wallet. 1 person can have many wallets. For each wallet you create, you need to write down a separate 24-word recovery phrase, set a separate PIN, and activate separate 2FA methods.

Crypto wallets store your private keys, keeping your crypto safe and accessible. They also allow you to send, receive, and spend cryptocurrencies like Bitcoin and Ethereum.

AvalancheGo Node provides a keystore API to store users Keys(Private-Public) but it is just for the testing only. Keys should be stored in frontend, not in the Node, which can be done using AvalancheJs. Always use AvalancheJS library to communicate with AvalancheNode. Don't get confused with `AvalancheGo`'s Keystore API and `AvalancheJs`'s Keystore API, Since AvalancheJs's Keystore API is built in client side so safer. `https://wallet.avax.network/` is an example of this library.

## UTXO Model (TRransaction vs Block)

Outputs of transactions are used as inputs of future transactions.

## Minting

Minting something means you create it. Minting a coin means you created that coin.

## QA

Q. Hi All, Hope you are doing well. I am new to blockchain technology. I have fundamental knowledge of how a blockchain work. But I have a confusion for building Avalanche Custom Blockchain. Let's say I built a custom VM and so its instance will a custom BlockChain. Now, Who is going to validate my Blockchain? since it needs the custom VM so will the validators on the Avalanche Primary Network will validate it? Thanks.  
A. You will need a good incentive to lure existing validators for validating your blockchain or a good chunk of AVAX to deploy your own avalanche validators.

Q. so, lets say my Node becomes a validator and my node is now part of the Avalanche primary network . And I have my custom VM with my node. And I defined a public subnet. Now what will be the process to tell other validators to join my subnet?  
A. Every node would see the transaction on the P-Chain that created your subnet, so every node knows about every subnet. To actually get validators to join your subnet - that is up to you, the subnet owner. You need to advertise and incentivise validators. There’s a validator marketplace on Fujinet. Don’t want to shill here. DM me if you want to learn more about it.

Q. So What I was thinking that, I would start one AvalancheGo node locally(no bootstrapping from MainNet) and define a subnet in it. Now I will start another 6 AvalancheGo nodes locally(Bootstrapping done from the first Node) . Now the 6 node will join the defined Subnet. every Node will have my custom VM running and so can validate the custom blockchain. Is this possible?  
A. my own answer is If I do this that means all the 7 nodes are in my authorities and so I can change the data. So, it is same as centralized authority. But blockchain is the server of trust where server means multiple server and each server have their own authority. so hacking one server can be recovered. That's why people want to use MainNet so to make sure that it is a trusted blockchain.

Q. can I create a subnet without a validator.
A. https://docs.avax.network/build/tutorials/platform/subnets/subnet-faq/

## Conclusion

We can say that Blockchain network is a distributed network where we can store block and the network will validate/replicate and hence the network(nodes) needs transaction fee.

In avalanche, to maintain the trust of between your users and the server(AvalancheGo which will provide custom services to users) you must join the MainNet. To join the MainNet the node have to stake 2000AVAX minimum . Now, the duty of each node on the Mainnet is to maintain by validating 3 blockchains, which are P-chain(where node information are stores), C-chain(Etherium blockchain), X-Chain(store information about the validator stake amount,etc.). Now the owner of the node can create a subnet where it can be defined which node can join the subnet. the mainnet is public so everyhting is visible like how many subnet, how many nodes etc. But, private subnet is where predefined node can join. Now the owner of the node can build a custom blockchain by making a custom VM and provide the VM to the defined nodes in the subnet so they can validate ie. maintain the security of the blockchain. And those node can join the subnet and start processng the custom chain. To recruit validators to join your subnet the node owner must contact the owner of those nodes.

without joining the MainNet one can create subnet locally and starts developing a custom blockchain and test in on Fujinet.

to make sure that you wont hamper user data or to build the trust of your user.
A server with rest api will have a database, Now the database must be shared with multiple nodes. even if you hamper others will save. so, basically same instance of the server are at different places owned by different people. But to which server the REST Api should be called?? The answer is use 1 NodeBalancer server. but when creating user client software should request to the Node.

## REFERENCES

1. [Avalanche tools](https://docs.avax.network/build/tools/)
2. [Avalanche Network Explorer](https://explorer-xp.avax.network/validators)
3. [Avalanche Node detail](https://avascan.info/staking/validators)
4. [Avalanche APIs](https://docs.avax.network/build/avalanchego-apis/)
5. [Avalanche FAQ](https://support.avax.network/en/)
6. [Avalanche Wallet](https://wallet.avax.network/)
7. [Avalanche Indexer](https://docs.avax.network/build/avalanchego-apis/index-api)
8. [Avalanche KeyStore](https://docs.avax.network/build/avalanchego-apis/keystore)
9. [Avalanche recover public key from address](https://docs.avax.network/build/tutorials/tutorials-contest/red-dev-sig-verify-tutorial/#recover-the-signer-x-chain-address-in-dapp)
10. [cb58](https://github.com/moreati/cb58ref)
11. [Bech32]
12. [X-Chain Transaction Format](https://docs.avax.network/build/references/avm-transaction-serialization/)
13. [C-Chain Transaction Format](https://docs.avax.network/build/references/coreth-atomic-transaction-serialization)
14. [P-Chain Transaction Format](https://docs.avax.network/build/references/platform-transaction-serialization)
15. Discord channel use Market
16. https://crypto.bi/avax-validation/
17. [proper setup of node](https://github.com/ava-labs/mastering-avalanche/blob/main/chapter_09.md)
18. [a full transaction](https://klmoney.wordpress.com/bitcoin-dissecting-transactions-part-2-building-a-transaction-by-hand/)
