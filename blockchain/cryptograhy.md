## Algorithm

- Cryptography

  - Asymmetric VS Symmetric  
    in Symmetric only one key used for encrypt and decrypt. Example algorithms - AES,BLOWFISH,  
    in Asymmetric two key used for encypt and decrypt. Example algorithm - RSA, DSA

  - Hash SHA256 (Secure Hash Algorithm is a cipher) is one flavor of SHA-2.  
    An length of data can be transformed into 256 bits. It is impossible to transformed back to original data from the hashed one. for a string A there will be a hashed X always and this mapping does not change.
  - **RSA**(Rivest-Shamir-Adleman) which is an **asymmetric cryptography algorithm** which use this private and public key concept.

    ```
    consider two prime numbers say
    p = 12
    q = 17
    n = p * q
    theta(n) = (p-1) * (q-1)
    consider d as public key such that d and theta(n) have no common factors
    consider e as private key such that (e*d) mod theta(n) = 1
    Encrypted = data^d*mod(n)
    Decrypted = encrypted^e*mod(n)

    Application 1: in Encryption and Decryption:
    Data = 88
    Public Key  = 7
    Private Key = 23
    n = 187
    Encrypted   = 88^7 mod 187
    decrypted   = Encrypted^23 mod 187

    Application 2: in Digital Signature
    hashed_message = hash(Message)
    signature      = hashed_message^private_key*mod(n)

    verified_signature = signature^public_key*mod(n)
    check if hashed_message == verified_signature
    ```

  - DH and DSA also types of **asymmetric cryptography algorithm**
  - SSL/TLS  
     Certificate Signin Request sent by a computer to CA(Cerificate Authority), and CA will verify the request and then it will Sign it(using CA's Private key) and send back the certificate. Now the server who wants tls/ssl will install it. Now, it will present this certificate to client first and then client(browser) will verify its signature using the public key which will be obtained from that certificate by reading it, and then again it will read the certificate of the intermediate CA and do this till the browser gets the final certificate which public key is stored in browser which is called ROOT PROGRAMS and now it can verify. So CA can be an intermediate CA which was verified by another CA. Types of SSL/TLS

        - Domain Validation SSL certificates – Assert server identity
        - Organization Validation SSL certificates – Assert partial organization identity
        - Extended Validation SSL certificates – Assert complete organization identity

    After the browser verifies the certificate present by the server using certificate chain, it's timw to handshake. Both server and browser will be using a cipher suite where a group of crypto and cipher algorithm will be present and this information will be presented inside the SSL/TLS certificate. so there will be list of cipher suite and client have to choose one. Now client will generate a symmetric session key and encrypt it using the public key and send it to the server and then server will be able to decrypt it since it has the private key. and that's what ssl/tls is, we shared a symmetric private key using asymmetric fashion.

    1. FAQ 1: HOW CLIENT WILL ASK THE SERVER FOR SSL/TLS CERTIFICATE? IS IT A REST GET/POST REQUEST?
       First of all it is at between transport and application layer. SO, HTTP->SSL/TLS->TCP->IP->FRAME, and so client can't be a HTTP client so no REST. it must be a SSL/TLS client. and browser have it.
       for example: openssl client
       ```
       openssl s_client google.com:443
       ```
    2. FAQ 2: WHAT'S IN THE CERTIFICATE AND HOW IT LOOKS LIKE? and IS IT JUST THE CERTIFICATE ONLY? OR DOES THE SERVER SEND OTHER INFO ALONG WITH IT?
       The certificate is a PEM file. it has label and data inside it. the data is a base64 encoded string which decode information is a DER/BER encoded data and which decoded information is serialized form of **X.500** data structure.
       so, certificate_info = b64(ber(asn.1 x.500))

       ```
       -----BEGIN CERTIFICATE REQUEST----- and -----END CERTIFICATE REQUEST----- show a CSR in PEM format.
       -----BEGIN RSA PRIVATE KEY----- and -----END RSA PRIVATE KEY----- show a private key in PEM format.
       -----BEGIN CERTIFICATE----- and -----END CERTIFICATE----- show a certificate file in PEM format.
       -----BEGIN PUBLIC KEY----- and  -----END PUBLIC KEY----- show a public key in PEM format

       ```

       **ASN.1** is a standard **interface description language** for defining data structure that can be serialized/deserialized in a cross platform way. it is like JSON,PROTOBUFF etc. so, it has a compiler. see wikipedia for list of different IDL.
       **X.500** is a series of data structure defined by International Telecommunication Union, where as **X.509** is one of them and is made by ASN.1. so **X.509** is just a data structure made by ASN.1. SO, to read it we need to deserialize using ASN.1 X.509 deserializer. and SSL/TLS client have it.

       ```

       openssl x509 -in certificate.crt -text -noout

       ```

       the **X.509** data structure contains this information -

       ```

       Certificate
          Version Number
          Serial Number
          Signature Algorithm ID
          Issuer Name
          Validity period
              Not Before
              Not After
          Subject name
          Subject Public Key Info # certificate public key
              Public Key Algorithm
              Subject Public Key
          Issuer Unique Identifier (optional)
          Subject Unique Identifier (optional)
          Extensions (optional)
          ...
       Certificate Signature Algorithm #the algorithm used to sign for example sha256WithRSAEncryption
       Certificate Signature #this is the certificate signature
       ```

       How this certificate signature is verified?
       Now we have the Certificate data, Certificate Signature and we can get the public key from certificate data.

       ```
        openssl x509 -pubkey -noout -in cert.pem > pubkey.pem
       ```

       cert.pem is the file where this ---BEGIN---.....END--- present. the output of the above command will be -

       ```
        -----BEGIN PUBLIC KEY-----
        MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAELE29VvMBi6FCjntEw89suGSGNxM4
        2LqH51uyaHIyet5T6aNDA8WmWB/t2WIgsuDn5q+zKf1goiQ9eBgyzpZ2LQ==
        -----END PUBLIC KEY-----

       ```

       Now from where CA authority put the Public key in that certificate? The answer is **CSR** which was also a PEM file send to CA authority which includes information about the server information and the public key and etc. we can make CSR using openssl. and then we have to copy the content of the pem file and put in a CA.

    3. FAQ 3: What's the difference between SSL and TLS?  
       Well TLS 1.0 is just upgraded version of SSL 3.0 which had security flaws.
       > So, Conclusion is server makes a CSR and is sent to CA and CA send the CERTIFICATE to server and server installs it so whenever ssl/tls client request it, it will provide, and then client ssl will make a symmetric key and encrypt it with public key(gets from the certificate) and send to server and since server has the private key at the first place so it will decrypt the key and get the symmetric key which was made by the clent. Now, since symmetric key is used for both encrypt and decrypt so both parties will be able to communicate. the .pem file format and DER/BER and x509 is useful to do all the core work.

- Block

  ```
  Information{
    prevhash
    nonce
    dificulty_level
    data
    hash
  }
  The hash is calculated using hash(prevhash,nonce,difficulty_level,data) but the condition is first 4 LSB of the hash must be 0. and to achieve this the nonce value can be varied.
  In Data, transaction will be signed using signature. which can be verified using its associated public key.
  ```

- Consensus Algorithm
  > Validators == miners, In PoW system, A validator must solve the hashing challange so after that it can verify a transaction.
  - Nokomoto consensus (ProofWork)
  - Avalanche Consensus (BY Team Rockey, 2018)
    - AVALANCHE IS A PLATFORM OF PLATFORM (say, bitcoin and etherium can communicate, it's a blockchain agnostic tech)
    - AVM (Avalanche Virtual Machine)

## Virtual Machine

> A software called client `VM` which follows a protocol/specification `PX` defined by `X` and is implemented by a programming language `P0` and that software takes input some Bytes called `BX` which is serialized/formatted and produced by a another software written in programming language `CX` as defined by `X` and executes Machine Instruction/some functions by the mapping of `BX` to Machine Language/functions/some specific work.

for example, A Java Client which includes -

```
client VM = HotSpot
P0 = c/c++
PX = JVM Protocol
X  = Oracle
BX = Java Byte Code
CX = C/C++
```

an Etherium client which includes -

```
client VM = Py-EVM/evmone/etheriumjs-vm/eEVM(P0=C++)/Hyperledger-Burrow(P0=G0)
```

So, A `Execution Client` includes a VM which executes ByteCode.
for example, Etherium `ETH1` which is `Execution Clinet` implementations are -

```
Geth/Erigon(Go)
Resu(Java)
Nethermind(C#,.NET)

```

`ETH2` which is a `Consensus Client` implementations are -

```
Prysm(GO)
Teku(Java)
```

if a Device is running either an `ETH1` or `ETH2` then it's called a `Node`. A `Node` Provide `REST` Api to the outer world to communicate.

Again there are different types of Node types -

```
FUll Node (Miners)
   Full Node stores the entire blockchain, so, basically 350GB or 160GB
Light Node.
   This does only transaction. only stores blockheader. it requests to full nodes.
Archive Node
   FUll Node + Historical Data.

```

Networks

```
Private Networks of Nodes:
   Development Networks
   Consortium Networks
Public Networks of Nodes
   MainNet Netoworks
   TestNet Networks

```

## Blockchain - A Digital trust

1991 : Stuart Haber in his research paper- "How to digitally timestamp a document"

The document/transaction stored in a block and then it is broadcast to all nodes. But, only initiator of transaction can decrypt the transaction.

## REFERENCES

1. youtube for graphical representation of blockchain
2. IDL [https://en.wikipedia.org/wiki/Interface_description_language]
3. WORLD STANDARD ORGANIZATION
   1. International Telecommunication Unit
   2. International Organization for Standardization
   3. International ElectroTechnical Commission
4. OPEN STANDARD ORGANIZATION
   1. Internet Engineering Task Force
5. International Standard Organization
   1. World Wide Web Consortium
6. American National Standard Institute
7. IEEE
8. OASIS
9. try google "international standard for \_\_\_\_"
10. [all Etherium Nodes](etherscan.io)
