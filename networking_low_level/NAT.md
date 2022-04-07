https://medium.com/@shivamrawat_756/guide-to-p2p-over-internet-9f7cf41470bd
https://tools.ietf.org/html/rfc5128
https://resources.infosecinstitute.com/topic/udp-hole-punching/

https://webrtc.github.io/samples/src/content/peerconnection/trickle-ice/
turnutils_uclient -T -W XXXX turn.example.com (pacman -S coturn)  or use turnutils_stunclient
https://manpages.debian.org/testing/coturn/turnutils_uclient.1.en.html
https://anyconnect.com/stun-turn-ice/



STUN CLIENT:
[bipul@archlinux ~]$ turnutils_stunclient -p 19302 172.253.121.127
0: : IPv4. UDP reflexive addr: 106.197.12.56:36579
0: : IPv4. UDP reflexive addr: 106.197.12.56:36579



IPV4 uses NAT specifically NAPT
	    Because your router don't know what to do with the request. When a request comes the router receives it but don't know what to do so it discards it. 
	    You need to tell your router what to do with the packets . using-
		1. port forwarding in router settings
		2. use ipv6
		3. STUN TURN ICE
			STUN servers connect to the user and returns their external IP address and external PORT (similar to HOLE PUNCHING)
			bittorrent use UDP HOLE PUNCHING.
				if SYMMETRIC NAT, use MULTIPLE HOLE PUNCHING
				so TURN server used to connect with symmetric NAT
			


UDP HOLE PUNCHING
our goal is to add the connection to the NAT table
	local 	  pc: nc -u -l -p 12001
	local     pc: sudo hping3 -c 1 -2 -s 12001 -p 8888 103.142.175.221
	public ip pc: nc -p 8888 -u 106.197.12.56 12001
	port collision may happen, so change the port



BOTH BEHIND NAT
	PC 1. sudo hping3 -c 1 -2 -s 33333 -p 44444 35.236.133.171 //will show packet loss
	PC 2. sudo hping3 -c 1 -2 -s 44444 -p 33333 106.203.146.232 //will show packet recv
	PC 2. nc -u -l -p 44444
	PC 1. nc -p 33333 -u 35.236.133.171 44444
	   AND BOOOOM it works on anywhere		
			OR
	PC 1. nc -p 12001 -u 49.37.34.248 23456
	PC 2. nc -p 23456 -u 35.189.175.200 12001
			OR
			
	iff ask stun
	./stun 74.125.197.127 19302 22222
	PC 1. nc -p 12345 -u 223.238.125.111  23456
	PC 2. nc -p 23456 -u 223.238.125.111  12345	
	
	nc -p --- -u `curl -s ifconfig.me` `./stun 74.125.197.127 19302 23456 | awk -F':' 'FNR == 7 {print $3}'`



A. endpoint mapping			: PRIVATE<-->PUBLIC_NAT<-->PUBLIC_SERVER
B. endpoint-independent maping		: resues the port mapping(for same internal ip and port) to any  external ip and any port
C. Address-Dependent Mapping		: reuses the port mapping(for same internal ip and port) to same external ip and any port
D. Address and Port-Dependent Mapping 	: reuses the port mapping(for same internal ip and port) to same external ip and same port while is still active
E. Endpoint-Dependent Mapping		: B C D
F. Endpoint-Independent Filtering	: -filters out only packets not destined to the internal address and port X:x, regardless of the external IP address and 
port 
					   source (Z:z)
					  -The NAT forwards any packets destined to X:x
					  -In other words, sending packets from the internal side of the NAT to any external IP address is sufficient to allow any
					   packets back to the internal endpoint.
G. Address-Dependent Filtering		: -The NAT filters out packets not destined to the internal address X:x.  Additionally, the NAT will filter out packets from
 					   Y:y destined for the internal endpoint X:x if X:x has not sent packets to Y:any previously 
 					  -In other words, for receiving packets from a specific external endpoint, it is necessary for the internal endpoint to 
 					  send packets first to that specific external endpoint's IP address.
H. Address and Port-Dependent Filtering : -The NAT filters out packets not destined for the internal address X:x.  Additionally, the NAT will filter out packets 
					   from Y:y destined for the internal endpoint X:x if X:x has not sent packets to Y:y previously.
					  -In other words, for receiving packets from a specific external endpoint, it is necessary for the internal endpoint to
					   send packets first to that external endpoint's IP address and port.
I. Endpoint-Dependent Filtering		: G H



J. NAT-Friendly P2P Application			: by using a publicly addressable rendezvous server for registration and peer discovery purposes.
K. Endpoint-Independent Mapping NAT (EIM-NAT)  	: is a NAT device employing Endpoint-Independent Mapping,eg: BEHAVE-compliant NAT devices
L. Hairpinning					: If two hosts (called X1 and X2) are behind the same NAT and exchanging traffic, the NAT may allocate an address on 							  the outside of the NAT for X2, called X2':x2'.  If X1 sends traffic	to X2':x2', it goes to the NAT, which must
						  relay the traffic from X1 to X2.



 Techniques Used by P2P Applications to Traverse NATs
 1. relaying
 2. connection reversal
 3. works only when one of the peers is behind a NAT device and the other is not. 
 4. UDP Hole Punching: relies on the properties of EIM-NATs, even when both communicating hosts lie behind NAT
		       devices.
					Terminology:

					    A : Node 1
					    B : Node 2
					    S : Rendezvous Server

					Steps:

					    1. A send  S a requests connection to B.
					    2. S sends A’s address to B and B’s address to A.
					    3. A sends garbage message to B and B sends garbage message to A. (Both Get discarded by their respective NATs)
					    4. Step 3 is repeated.
					    Connection Established.
					    
					    
					    
					    
					    
					    
					    
		network traversal service	network traversal service	network traversal service	network traversal service		    
					    
1. A & B Always Listen to a static server S (S-> A's phone number + A's ext IP:PORT
					      -> B's phone number + B's ext IP:PORT)
2. A tells S to send A's ext IP:PORT to B and need B's ext IP:PORT
3. A calls B and B calls A using their IP:PORT
