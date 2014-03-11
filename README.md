

Darknet Plan
============

The Skycoin Darknet is a high performance, privacy preserving routing protocol inspired by cjdns. Users receive skycoin for contributing resources to the network and expend resources for using network resources.

The protocol is designed to operate over legacy internet and nodes physically connected by wifi. The long term objective is to create long distance point-to-point wifi connections which bypass existing internet providers.

User Stories: Link Aggregation
===============================

Bob has a 2 Mb/s internet connection. It takes Bob minutes to load Youtube videos. Bob and has five neighbors with 2 Mb/s connections. They install Skycoin nodes. Bob's Skycoin node connects to his neighbors through wifi and aggregates the bandwidth, giving Bob a 12 Mb/s connection.

Bob receives Skycoin for relaying traffic and expends Skycoin for using network resources.

Notes:
- Bob's IPv4 traffic tunneled over the Darknet enters the normal internet at a local server on a network backbone
- Bob's Skycoin connection appears as a VPN connection on his computer
- Bob's traffic may take multiple routes between his home and the IPv4 Gateway node
- With $150 in equipment, Bob can connect to nodes up to five miles away at 40 Mb/s
- With $1500 in equipment, Bob can connect to nodes up to 15 miles away at 1.4 Gb/s

User Stories: Backhaul
======================

Alice lives in a large city, 2 miles from a colocation center with terabytes per second of fiber optic backbone.

Alice's internet speed is 2 Mb/s.

Alice had cheaper, faster internet, before the FCC stuck down the common carrier access rules. Now, Alice only has one choice for internet. Alice's national ISP is the only ISP after the merger of the two largest cable companies in America.

After the merger the CFO stated "people dont want faster internet", raised prices and put in place bandwidth caps. Alice pays $0.30 per GB for going over her 100 GB bandwidth cap.

Alice's ISP has been getting worse after net neutrality was struck down by a secret backdoor international trade agreement, that even members of congress were not allowed to see or vote on before it was signed.

Alice's Youtube and Netflix videos are loading slower than ever before. Alice's ISP has started throttling Netflix, Youtube and Bitorrent while publicly denying it. 

Alice's ISP has begun tracking every website she visits, recording her personal information and selling it to the NSA and marketing companies. Alice's ISP has been stealing revenue from the websites Alice visits, by replacing the website's ads with its own advertisements. Alice's ISP is starting to blacklist websites it doesnt like.

Alice hears about Skycoin, finds another Skycoin user with an office in the colocation center. Alice pays $1500 and installs 1.4 Gb/second Ubiquity airFiber antenna on her roof to bridge the distance between her and the fiber backbone.  Alice's connection acts as the backhaul for her neighborhood's local Skycoin mesh.

Alice cancels her internet service.

User Story: Internet Kill Switch
=================================

<todo>

Hardware
========

For development, we are using the following

Platforms:
- Debian
- Raspberry PI / Beagle Boards

Wifi Devices:
- Edimax EW-7811Un (good linux support)
- TP-LINK TL-WN722N (external antenna support for long distance directional links)

Directional Antenna:
- TP-LINK TL-ANT2424B 24dBi 60 cm Directional Grid Parabolic Antenna. Up to 10 mile range for line of sight.
- see: http://fabfi.fabfolk.com/

Future:
- RONJA (see http://en.wikipedia.org/wiki/RONJA )
- HackRF http://kck.st/1eb5z2R
- Li-Fi http://en.wikipedia.org/wiki/Li-Fi

Technical Objectives
====================

Implementation:
- prototype in Golang
- extremely simple. less than 2000 lines of code
- minimal number of dependencies

Design Goals:
- Privacy preserving
- coin incentives for provisioning bandwidth, storage and backhaul
- Open Access Wifi mesh networks
- Designed to bridge last mile between the network backbone and home
- runs on Rasberry Pi and Ubiquity Hardware
- "zeroconf". Plug in and runs, no configuration
- difficult to detect and throttle

Technical Aspects:
- uses pubkey hashes as network addresses
- Link Aggregation (ability to aggregate bandwidth from multiple connections)
- instant, low overhead, distributed bandwidth micropayments using off blockchain transactions
- link layer, does not define routing
- store and forward?
- compatibility bridge with IPv6 networks

Security:
- Link layer encryption between nodes
- End to End Encryption
- Post Public Key encryption primitives

Protocol Draft
==============

Here is whole protocol.

The protocol is
- easy to implement in C
- simple (no backdoors)
- fast
- secure
- use the minimum number of cryptographic primitives

Link Layer:
- Open TCP socket to remote host. You need their pubkey
- First packet: encode session key as point on a secp256k1 elliptic curve using peer's public key. encrypt your pubkey with ChaCha20 using the session key. Send packet
- Response Packet: destination encodes second session key with your public key and sends response packet. This wrapper is encrypted with the ChaCha20 key you sent previously.
- You now have bidirectional connection to node for sending and receiving data
- there is seperate session key for each direction of communication

Routing:
- At each node, a "path" is established. For each node in a route, you register the next node and receive a 32 bit int. The 32 bit int when prefixed on packet determines the node packet will be forwarded to.
- Each node decodes the packet and pops off first 4 bytes to determine next node to transport packet to.

Payment for Transport:
- Nodes keep track of how much traffic goes each way for the route
- The person intiating the route makes an escrowed "confidence payment" with a 120 byte off block chain Skycoin Transaction.
- The origin node clears payment with the node every few minutes

Note:
- route is determined by origin of traffic
- the destination can communicate back to origin but cannot identify origin node
- payment overhead is 120 bytes per payment
- per hop overhead is 4 bytes (exercise for reader: make it constant)
- public keys are never exposed as plaintext in protocol
- cannot communicate with node without node public key
- 32 bit route path prefix information should be obfuscated by shared secret with node
- packets should be fixed length or multiple of power of 2 for secure applications to resist traffic analysis
- the pubkey a node is sending from can be thrown away. Destination pubkey hash acts as network address for routing. Destination pubkey only decrypts, never signs. Sucessful decryption of session key is proof of private key possession and identity.

Todo:
- this is transport layer protocol. protocol layer over this layer sends traffic over multiple paths to the destination, using fountain coding.
- since origin determines path, origin can optimize for latency or throughput and other criteria
- traffic and handshake should be disguised as SSL protocol to deter throttling by ISPs

Privacy
=======

A user operating a Skycoin Wifi access point allows any user in range to connect through that access point. The access point operator cannot determine the nature of the traffic passing through the access point because it is encrypted. The recipient of the traffic is unable to determine that the path of the traffic passed through the access point.

This effectively removes legal liability for operating public access points. The operator neither has any information about the traffic being relayed nor can the recipient of traffic identify the operator of the network entry point.

Furthermore, with the addition of a mandatory hop (a "guard node") it is impossible for ISPs to easily identify that Skycoin Darknet traffic from a particular public access point has been relayed through a particular cable modem.

Summary:
- Skycoin Darknet Wifi access points are public by default
- Access point operators cannot see contents of traffic through the access point
- Access point operators cannot see destination of traffic routed through the access point
- The recipient of traffic cannot determine the origin or path the data traveled through the network
- Using "guard nodes" ISPs cannot determine that traffic from a particular access point is being relayed through a particular terminating connection (cable modem)

Tradeoffs
=========

The Skycoin Darknet protocol is low latency, high throughput and offers a greater degree of privacy than previous systems. However to achieve these goals, several tradeoffs were necessary.

1. Routing decisions are pushed to the origin node instead of the network
2. Rasberry PIs can only forward 150 Mb/s of second of traffic due to encryption overhead. FGPA hardware could accelerate this to GB/s.
3. The network has the best performance and lowest overhead for large, latency insensitive transfers. Torrents will do very very well over network.
4. Real time applications sending many small packets will function over network, but incur larger overhead than TCP/IP. The theoretical latency and "jitter" in latency is lower than for TCP/IP, but with higher bandwidth requirements for real time applications such as VOIP.
5. Bandwidth microtransaction pollute the blockchain. Therefore we are relying on trusted third parties for low overhead off-blockchain microtransactions for bandwidth confidence payments.
6. Since routing decisions are pushed back to the origin client, clients must maintain routing tables or rely upon 3rd party servers for routing information.
7. Store and Forward operation increases network throughput and reliability, but degrades quality of service for real time applications. Store and Forward operation introduces additional ram and storage requirements on nodes, which may tax the capacity of traditional routers.
8. Nodes that interface between the Skycoin Darknet and traditional internet may suffer the same problems as Tor exit nodes. Most tor exit nodes are blacklisted for editing wikipedia or creating user accounts on websites due to spam issues. To maintain a high quality of service, exit nodes may require trust relationships, payment or user registration to prevent abuse.
