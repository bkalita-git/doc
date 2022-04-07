## HTTP

HTTP has four versions — HTTP/0.9, HTTP/1.0, HTTP/1.1, and HTTP/2.0.

0. How HTTP works:  
   It's all about socket. for both TCP and UDP socket the connection has to established first and then can communicate. in HTTP for one GET request there may be several TCP packets. and that's the difference of working in HTTP layer. So, HTTP packet is converted to TCP packets and those will be sent to the server. and then close the TCP connection.

1. HTTP/1.0 over TCP
   open the TCP connection and send GET request and close the TCP connection when got its response. for one GET request there may several TCP packets since it breaks down things. For the next GET it has to again establish the connection to the server and then close it when got response.
   Streaming with buffer transfer.
2. HTTP/1.1 over TCP
   - invented Keep-alive header
   - Streaming with chunked transfer.
   - pipelining is disabled by default.  
     open the connection and sent GET request with keep-alive and and got its response and no need to close it. again client can send GET request since the TCP connection is already open.
3. HTTP/2.0 over TCP
   - /HTTP/2.0 needs HTTPS by default
   - compression
   - multiplexing
   - server push
   - SPDY
   - Protocol Negotiation during TLS (NPN/ALPN)
     - Next Protocol Negotiation
     - Application Layer Protocol Negotiation (extension of TLS)
4. HTTP/2.0 over QUIC(HTTP/3)
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
