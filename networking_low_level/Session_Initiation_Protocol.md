Application Layer Protocol
It is a signaling protocol
It is a text-based protocol
SIP carries SDP as SIP payload
SIP can be carry with TCP,UDP,SCTP
SIP can be encrypted with TLS
SDP payload typically employs the RTP or SRTP
SIP is used in video conferencing, streaming media distribution, instant messaging, presence information, file transfer, internet fax, online games
SIP jos is to setup and terminate the connection
for call setup the body of SIP message contains SDP which specify media format, codec, media communication protocol
voice and video streams are typically carried using RTP or SRTP
SIP resources are identified by a URI, typically sip:username@domainname or sip:username@hostport , hostport can be an ip address or domain+port, for secure transmission sips is used instead of sip

in SIP both user agents(client application) act as server(UAS) and client(UAC), UAC sends SIP REQUEST then UAS receives REQUEST and returns a SIP RESPONSE 
in SIP the UA may indentify itself using a message header field containing  a text description of the software ,hardware or the product name, it is send in REQUEST messages

## NETWORK ELEMENTS:
	1. User Agent
	2. Proxy Server
	3. Redirect Server
	4. Registrar ->(Located in)-> Proxy Server
	5. Session Border Controller
	6. Gateway
