## Payload

when the payload is large in a http request, then there will be many TCP packet created and all those info will be in the http request.
so, the size of http payload will be small since it will only contain the metadata of all tcp packet.

there will be 10 TCP Packet out of them 1 TCP packet will contain HTTP Packet where it will store LINK to other 9 tcp packet. and those other 9 tcp packet will contain the data.

## p2p Signaling

It is used to exchange each other public ip:port
