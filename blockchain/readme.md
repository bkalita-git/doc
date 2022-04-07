## Algorithm

- Cryptography

  - Hash SHA256 (Secure Hash Algorithm) is one flavor of SHA-2.  
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

## REFERENCES

1. youtube for graphical representation of blockchain
