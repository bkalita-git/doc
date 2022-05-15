## HTTP

HTTP has four versions — HTTP/0.9, HTTP/1.0, HTTP/1.1, and HTTP/2.0.

- TCP SLOW START to find out optimal bitrate -  
   for a tcp connection after tcp handshake is done, client will send data to the destination say 5 packets and it gets response of all 5 packets, next time the client will send 10 packets and it gets response of all 10 packets, then client will gradually increase the number of packets till their reponses are received correctly. so that's how optimal bitrate is found. So, obviously for multiple tcp connections this step has to be done for all, hence performance decreased to find the optimal bitrate. So, What http2.0 does is that always have 1 tcp connection. and for each data/packet it transmit it assigns a stream id and so response has the stream id too. we can send 5 packets for steam 1 and 2 packet for stream 2 like that. so, no order has to maintained and this is multiplexing.
- Server Push  
  after tcp handshake let's say the client send a GET request for a html file and then the server sends response for it, also the server sends some extra file like css, js file in 2 reponses without asked by the client. SO, the browser can cache this resources.
- Prefetching
  since server push is still experimental, so, in html file one can write

  ```
  <link rel="prefetch" href="banner.png">
  ```

  so, server can do the push and also client knows now that it will receive banner.png without asking the server.

- Hinting Approach

* Head of Line Blocking:  
  A queue containing http packets, where first packet is blocking others since it must be sent so other can proceed. in HTTP1.1 it occurs when the number of allowed parallel requests(typically 6 tcp connection) in the browser is used up, and subsequent requests need to wait for the former ones to complete. HTTP2 uses **multiplexing** which removes HOL but still present at Transport layer, in HTTP2 QUIC this problem also solved.  
  In HTTP2 using one connection multiple HTTP packet can be sent.
  However, HTTP1 have a concept of **Pipelining** where multiple requests sent at once, however they still had to return in order. But in HTTP2 order does not matter, response can arive at any order.

1. How HTTP works:  
   It's all about socket. for both TCP and UDP socket the connection has to established first and then can communicate. in HTTP for one GET request there may be several TCP packets. because tcp packet has fixed size so if the http packer is bigger than that it will get chucked to many tcp packets. and that's the difference of working in HTTP layer. So, HTTP packet is converted to TCP packets and those will be sent to the server. and then close the TCP connection.

2. HTTP/1.0 over TCP
   open the TCP connection and send GET request and close the TCP connection when got its response. for one GET request there may several TCP packets since it breaks down things. For the next GET it has to again establish the connection to the server and then close it when got response.
   Streaming with buffer transfer.

3. HTTP/1.1 over TCP
   - invented Keep-alive header
   - Streaming with chunked transfer.
   - pipelining is disabled by default.  
     open the connection and sent GET request with keep-alive and and got its response and no need to close it. again client can send GET request since the TCP connection is already open.
4. HTTP/2.0 over TCP

   - /HTTP/2.0 needs HTTPS by default
   - It Transfers binaries instead of text.
   - is fully multiplexed, instead of ordered and blocking
   - can therefore use one connection for parallelism
   - uses header compression to reduce overhead (HPACK)
   - allows servers to “push” responses proactively into client caches

   1. If it uses TLS/SSL then why do we need HPACK?  
      It's called **CRIME** attack. in http1 and http2 header are compressed too and then encrypted(in http2 only) using ssl/tls symmetric session key. SO, nobody will know what are the headers but as a CORS attack if some one tries to send a request behalf on you using cookies stored in your browser then it can send. But it does not know/read what the cookie is since it's compressed and encrypted. so what the attacker do is try to add random text into headers and while compression similar text have 1 copy only so the size will be same and the attacker can guess which text is present on that header. SO, HPACK is the solution.
      AGain, **BREACH** attack. the content/payload is compressed too when content-encoding:gzip.

5. HTTP/2.0 over QUIC(HTTP/3)
   - Replaces TCP with QUIC (UDP with congestion control)
   - All HTTP/2 features
   - Experimental

## Transloadit Upload Server (TUS) Protocol

1. for this the server creates the file or if already exist then returns the file address.

   ```
   POST /files HTTP/1.1
   Host: localhost:8080
   Content-Length: 0
   Upload-Length: 250

   HTTP/1.1 201 Created
   Location: localhost:8080/files/12
   ```

2. client sends a head request for patch to the path

   ```
   HEAD /files/12 HTTP/1.1
   Host: localhost:8080

   HTTP/1.1 200 OK
   Upload-Offset: 0
   ```

3. client uploads

   ```
   PATCH /files/12 HTTP/1.1
   Host: localhost:8080
   Content-Length: 250
   Upload-Offset: 0
   [0 to end of the file]
   ```

4. failed in uploads say 100 bytes then client proceeds from step 1 to 2 and

   ```
   HEAD /files/12 HTTP/1.1
   Host: localhost:8080

   HTTP/1.1 200 OK
   Upload-Offset: 100
   ```

5. uploading from 100 bytes

   ```
   PATCH /files/12 HTTP/1.1
   Host: localhost:8080
   Content-Length: 150
   Upload-Offset: 100

   [100 to end of the file]
   HTTP/1.1 204 No Content
   Upload-Offset: 250
   ```

## REFERENCES

1. Resource for IETF (Internet Engineering Task Force) [https://datatracker.ietf.org/]
