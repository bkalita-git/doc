## UDP AND TCP

The meaning of server and client. Multiple Socket can be associated with same IP:PORT if SO_REUSEADDR is used.

After the TCP Handshake any party can send packet to the other one before closing the connection. it's not like that client will send and server will response. in TCP, after handshake is done server can send packet to client too without client's request. After the handshake , client only has to listen to the connected stream for incoming packet since server will be able to send the packet to the already connected stream and vice versa. So, in TCP we first build the stream using TCP Handshake, After that both party will be able to communicate idependently with each other using the stream.

In UDP, we don't build any stream.

## Payload

when the payload is large in a http request, then there will be many TCP packet created and all those info will be in the http request.
so, the size of http payload will be small since it will only contain the metadata of all tcp packet.

there will be 10 TCP Packet out of them 1 TCP packet will contain HTTP Packet where it will store LINK to other 9 tcp packet. and those other 9 tcp packet will contain the data.

## p2p Signaling

It is used to exchange each other public ip:port
