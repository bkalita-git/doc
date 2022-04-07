## peer discovery

mDNS and DHT

## multicast DNS (mDNS) and a Distributed Hash Table (DHT).

mDNS doesn’t work over the internet and DHT does. DHT requires BOOTSTRAP NODE. Other nodes can join the network by connecting to a bootstrapping node, and then discover the rest of the network. the BOOTSTRAP NODE must be accessible from public.

libp2p AddrsFactory

NAT TRAVERSAL
https://github.com/orgs/libp2p/repositories

```
go-libp2p-core
    Routing/    : An interface
        ContentRouting : An interface
        PeerRouting    : An interface
        ValueStore     : An interface
    Discovery/  : An interface
        Advertise      : An interface
        FindPeers      : An interface

go-libp2p-kad-dht
    implements all those above routing

libp2p-dht-discovery
    RoutingDiscovery struct implements discovery interface
```

https://hackmd.io/@juincc/Hk_0PWFhS
