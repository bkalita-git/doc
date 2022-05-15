- [BlockChain 101](#blockchain-101)
- [Transaction](#transaction)
  - [P2PKH (Pay to Public Key Hash)](#p2pkh-pay-to-public-key-hash)
  - [P2SH (Pay to Script Hash)](#p2sh-pay-to-script-hash)
  - [P2WPKH (Paying to witness Public Key Hash)](#p2wpkh-paying-to-witness-public-key-hash)
  - [Coinbase Transaction](#coinbase-transaction)
  - [UTXO (Unspend Transaction Output)](#utxo-unspend-transaction-output)
  - [Avalanche New Blochain creation Process with 5 nodes](#avalanche-new-blochain-creation-process-with-5-nodes)
- [REFERENCE](#reference)

## BlockChain 101

The same thing done using blockchain technology is possible using client-server architecture, Though Blockchain technology uses distributed architecture. We can consider the Distributed Blockchain Network as having multiple server providing services/api and most of the server are own by different person. We can't directly say that Blockchain distributed network is p2p, Because User(person) will call Api to 1 server among the group of Server in that distributed network. While P2P like bittorrent software act as client and also server.

Consider having a server with a database attached with it, Now the owner of the Server can alter the data also if something happens to the server, the users will not be able use the service. But in distributed Blockchain network, There can be different Owner of those server and also the algorithm works is a way so that even if one owner alters Data still user will get the correct data. So basically Each server must have correct replica of Data.

But what does it mean by a User on the blockchain network? And How user stores Data on blockchain network? Who give permission to the user to allow APi call by server among the blockchain network? Well, A user on blockchain is not there till that user issues a transaction to one of the server and also the server should allow the user to use the API. So, Basically Transaction are built on client side and then send it to one of the server. But here is the catch, the information about the user are inside the Transaction and also the paying cost for processing the transaction is inside the Transaction. Paying Cost for processing the transaction? Yes, you have to pay to those server obviously.

Before proceeding to the transaction lets talk about the database. The database can be of "Blockchain" or "DAG" architecture. Blockchain is a linear List(chain) composed of Block and each block contains serveral Transactions(above we mentioned, where user issues to server). On the other hand, Dag is a DAG composed on vertices and each vertices is composed on one or more transactions. And may be there are other types of Architecture of the database but Bitcoin, Ethereum uses Blockchain and Avalanche uses Both Blockchain and DAG architecture. But, still these are called Blockchain Technology. Every components of Blockchain technology has ID. for example, the Database (a particular Blockchain/DAG) has ID, Vertices/Blocks also have ID, Transactions have ID, Users have ID(Address), all the running Server on the Blockchain Network has ID etc. Different types of ID are made differently. For example, User ID or specifically Address of the User can be made from User's Public Key which also derived from User's Private Key. On the Other hand, Transaction ID can be made by taking Hash Value of the Transaction Packet. etc. So, different Platform may slightly change the algorithm to generate ID of their components but it's important to study about it before develop application on a particular platform.

## Transaction

The core atomic information in blockchain network is a transaction. Consider it as a packet with ID. Different platform may have different fields in transaction packet. We will study some model of Transaction. Transactions are core component of the blockchain system because it holds everything. User builds transaction with the help of their Private Key. Before we even discuss about how a user builds a transaction packet and where the user is sending money to other users, We have to first understand the concept of Money in blockchain and from where it comes from. Well since blockchain database is composed of transactions in the end, There must be a transaction at the begining. Also, **a blockchain database does not exist without a transaction at first**. So we call it Genesis Data. And about the money from where it comes at the first place? Well it's created. Yes, Money is created at the first time and it's called digital asset. So, The first transaction is special one because it creates Money. and this transaction is called "CoinBase" Transaction.

Alright, lets talk about the Address of a user. Well, it can be P2PKH, P2SH, P2WPKH(or bech32)

So, How is it possible to issue the first transaction ever(coinbase) to the server and how server is going to get paid? Well, **The first transaction to build the blockchain dataset for the first time is to built/issued by the Server, Not the user and this is the actual genesis data (The first block and inside there is a coinbase transaction)**

Well! Well! Well! I said the first transaction is built by the miner/server and it creates money by just saying that It sent 50BTC to these of the addresses. And this is true for all the blockchain platform. For Avalanche network, The first transaction ever was with ID is `11111111111111111111111111111111LpoYY` and it had also no input but had 100 outputs which are just 100 addresses with money written by the miner, So basically it was a coinbase transaction and this transaction was in `Block -1` and the ID of that block was `UUvXi6j7QhVvgpbKM89MP5HdrxKm9CaJeHc187TsDNf8nZdLk`. Along with the coinbase, The block had another `77 transactions` out of where 2 transactions - transaction id `2q9e4r6Mu3U68nU1fYjgbR6JvwrRx36CohpAX5UQxse55x1Q5` was about type `CREATE CHAIN` `C-CHAIN` and transaction id `2oYMBNV4eNHyqk2fjjV5nVQLDbtmNJzq5s3qs3Lo6ftnC6FByM` was about type `CREATE CHAIN` `X-CHAIN`. Now the interesting thing is, These 3 transactions ID(Not the Block ID) are the ID of the `P-CHAIN`, `C-CHAIN` and `X-CHAIN` and also the ID of P-CHAIN is the `SubnetID`. Now Other 75 transactions were about type `ADD VALIDATOR` and in each of these transaction has multiple outputs but no input, These output were just defining money/AVAX to some platform addresses as stacking set and stacking reward. and This `Block -1` was the genesis data of Avalanche Blockchain, which basically was the `P-Chain`.
HMM, **SO when I start a node of avalanchego then it first creates the `Subnet` and genesis block where creation of `c-chain`,`x-chain`, and adding validators to it were defined but these genesis data varies with type of the network. for mainnet, local, testnet these are defined differently**. Also for localnet, the genesis contains a `ewoqjP7PxY4yr3iLTpLisriqt94hdyDFNgchSxGGztUrTXtNN` private key so it can be used to fund the local network.

### P2PKH (Pay to Public Key Hash)

### P2SH (Pay to Script Hash)

### P2WPKH (Paying to witness Public Key Hash)

### Coinbase Transaction

Now this time not the user but The server/miner builds a transaction. In this transaction it states that 50BTC is coming from nowhere and send to a BTC Address. Now, this BTC Address is the user who own this Server/miner. So, before running the server/miner software the owner puts his/her bitcoin address in to that. Now the server combines this coinbase transaction with other verified transactions and builds a block and does the consensus mechanism(PoW ie solving Nonce for perfect hash in case of bitcoin) for this and upon success the block gets added to the blockchain and so each server on the network will have this information.

This process of building coinbase transaction which states that 50BTC comes from nowhere and sent to an address is called **Minting coin** and also called **Block Subsidy**.

Together Block Subsidy(coinbase transaction) and Transactions Fee of all transactions in a block makes **Block Reward**

### UTXO (Unspend Transaction Output)

Lockscript and unlockscript are there, Inputs and outputs are there. changeAddr are there.

### Avalanche New Blochain creation Process with 5 nodes

0. using ANR make 5 local nodes
1. user1.js file which make a keystore and imports `ewoq` key on P-Chain and so it gets its P-Chain address. (need both avalanche js + avalanche wallet)
2. user1 Issues `Create Subnet` Transaction on P-Chain to any Node.
3. user1 Issues `add validators` transaction on P-Chain to any Node for all 5 local nodes. (so it stakes 2000AVAX\*5)
4. user1 Issues `Create Blockchain` Transaction on P-Chain to any node

## REFERENCE

1. [The first transaction on bitcoin](https://blockchain.info/tx/4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b?format=json)
2. [The first pchain avalanche block](https://avascan.info/blockchain/p/block/-1)
3. [Avalanche network bootstrapping locally](https://github.com/blockchain-private-network/avalanche-local-private-network/blob/main/start.sh)
4. [UTXO](https://docs.avax.network/specs/avm-transaction-serialization#utxo)
