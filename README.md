

Darknet Plan
============

The Skycoin darknet is designed for a world where net neutrality has failed. A world where monoplistic cable companies have become the gate keepers to the internet.

The Skycoin Darknet is a technological response to ACTA, PIPA, SOPA and the Comcast, Time Warner Merger. The proticol is designed to bridge the last mile between fiber and the home and eliminate reliance upon monopolistic ISPs.

The Skycoin darknet is designed specificly for deploying open access wifi mesh networks and community ISPs. The network operates over fixed position point-to-point wifi connections using commodity hardware and the legacy internet.

Users receive Skycoin for contributing resources to the network and expend Skycoin for using network resources.

Network deployment will begin by July.


Hardware
========

For development, we are using the following

Platforms:
- Debian
- Raspberry PI
- Beagle Boards

Wifi Devices:
- Edimax EW-7811Un (good linux support)
- TP-LINK TL-WN722N (external antenna support for long distance directional links)

Directional Antenna:
- TP-LINK TL-ANT2424B 24dBi 60 cm Directional Grid Parabolic Antenna
- see: http://fabfi.fabfolk.com/

Project Milestones
==================

1. Two Raspberry PIs running nodes are connected to two different cable modems. A third computer connects to the Raspberry PIs via wifi and is able to aggregate bandwidth over both connections.
2. First coins are exchanged for network traffic


Technical Objectives
====================

Implementation Details:
- prototype in Golang
- extremely simple. less than 2000 lines of code
- minimal number of dependencies

Design Goals:
- Open Access Wifi mesh networks
- coin incentives for provisioning bandwidth, storage and backhaul
- Designed to bridge last mile between the network backbone and home
- resistant to latency, high packet loss and low reliability connections
- runs on Rasberry Pi and Ubiquity Hardware
- "zeroconf". Plug in and runs, no configuration
- difficult to detect and throttle

Technical Aspects:
- uses pubkey hashes as network addresses
- instant, low overhead, distributed bandwidth micropayments using off blockchain transactions
- link layer, does not define routing
- store and forward?
- compatibility bridge with IPv6 networks
- Link Aggregation (ability to aggregate bandwidth from multiple connections)

Security:
- Link layer encryption between nodes
- End to End Encryption
- Post Public Key encryption primitives

Privacy
=======

A user operating a Skycoin Wifi access point allows any user in range to connect through that access point. The access point operator cannot determine the nature of the traffic passing through the access point because it is encrypted. Furthermore the recipient of the traffic is unable to determine that the path of the traffic passed through the access point.

This effectively removes legal liability for operating public access points. The operator neither has any information about the traffic being relayed nor can the recipient of traffic identify the operator of the network entry point.

Furthermore, with the addition of a mandatory hop (a "guard node") it is impossible for ISPs to easily identify that Skycoin Darknet traffic from a particular public access point has been relayed through a particular cable modem.

Summary:
- Skycoin Darknet Wifi access points are public by default
- Access point operators cannot see contents of traffic routed through the access point
- Access point operators cannot see the destination of traffic routed through the access point
- The recipient of traffic cannot determine the origin or path the data traveled through the network
- Using "guard nodes" ISPs cannot determine that traffic from a particular access point is being relayed through a particular terminating connection (cable modem)


Q&A
===

Why not Bitcoin? Why build the network on a new coin?

We see no advantage in developing the darknet routing proticol on a Bitcoin payment system and there are several disadvantages compared to emerging alternatives.

- Bitcoin does not offer developers and early adapters the economic rewards of a new coin.
- Bitcoin cannot fund the required infrastructure investments. The Bitcoin community has been fundamentally unable to fund critical infrastructure projects. Infrastructure projects such as coinjoin have been absymal failures at raising funds for development. The project must be self funding to succeed and cannot rely upon donations from the Bitcoin community.
- Bitcoin has numerous security issues such as transaction malleability and hash collisions in coinbase outputs which fundamentally cannot be fixed or have been monkey patched. Skycoin fixes all known security and usability issues in Bitcoin.
- The Bitcoin blockchain is bloated. The blockchain size is too large for embedded devices and growing exponentially. When the Bitcoin blocksize limit is lifted, Bitcoin may become unusable for average users. Skycoin supports unspent output snapshots and takes measures to ensures clients will never require more than 1 GB of storage or RAM to operate.
- Bitcoin does not support off-blockchain transactions required for bandwidth confidence micropayment. Skycoin supports instant, free, low overhead, local microtransactions that do not pollute the blockchain. This is expecially important for maintaining network operation in environments where global blockchain consensus is disrupted by lose of global internet connectivity in target countries.
- Bitcoin fees are too high for bandwidth micropayments. Bitcoin fees are currently around 10 cents per transaction and increasing. Skycoin payments are "free", with fees paid in coinhours to limit spam.
- Bitcoin payments are too slow for bandwidth microtransactions. Bitcoin payments require 10 minutes to clear. Skycoin payments are 40 times faster.
- Several emerging problems are throwing the long term survival of Bitcoin into question. Bitcoin was a proof-of-concept and an experiment. Satoshi never believed that it would reach as far as it has without being superceded. Bitcoin was a seed of an idea and never intended as its final form.

Tradeoffs
=========

The Skycoin Darknet proticol is low latency, high throughput and offers a greater degree of privacy than previous systems. However to achieve these goals, several tradeoffs were necessary.

1. Routing decisions are pushed to the origin node instead of the network
2. Rasberry PIs can only forward 150 Mb/s of second of traffic due to encryption overhead. FGPA hardware could accelerate this to GB/s.
3. The network has the best performance and lowest overhead for large, latency insensitive transfers. Torrents will do very very well over network.
4. Real time applications sending many small packets will function over network, but incur larger overhead than TCP/IP. The theoretical latency and "jitter" in latency is lower than for TCP/IP, but with higher bandwidth requirements for real time applications such as VOIP.
5. Bandwidth microtransaction pollute the blockchain. Therefore we are relying on trusted third parties for low overhead off-blockchain microtransactions for bandwidth confidence payments.
6. Since routing decisions are pushed back to the origin client, clients must maintain routing tables or rely upon 3rd party servers for routing information.
7. Store and Forward operation increases network throughput and reliability, but degrades quality of service for real time applications. Store and Forward operation introduces additional ram and storage requirements on nodes, which may tax the capacity of traditional routers.
8. Nodes that interface between the Skycoin Darknet and traditional internet may suffer the same problems as Tor exit nodes. Most tor exit nodes are blacklisted for editing wikipedia or creating user accounts on websites due to spam issues. To maintain a high quality of service, exit nodes may require trust relationships, payment or user registration to prevent abuse.

User Stories: Link Aggregation
===============================

Bob has a 2 Mb/s internet connection and it takes minutes ot load Youtube videos. Bob and his five neighbors with 2 Mb/s connections install Skycoin nodes. Bob's Skycoin node aggregates bandwidth across the five neighboring nodes and acts as a 12 Mb/s connection.

Bob receives Skycoin for relaying traffic and expends Skycoin for using network resources.

Notes:
- Bob's IPv4 traffic tunneled over the Darknet enters the normal internet at a local server on a network backbone
- Bob's Skycoin connection appears as a VPN connection on his computer
- Bob's traffic may take multiple routes between his home and the IPv4 Gateway node

User Stories: Backhaul
======================

Alice lives in a large city, 2 miles from a colocation center with terabytes per second of fiber optic backbone.

Alice's internet speed is 2 Mb/s. Alice had cheaper, faster internet, before the FCC stuck down the common carrier access rules. Now, Alice only has one choice for internet. Alice's national ISP is the only ISP after the merger of the two largest cable companies in America. 

After the merger the CFO stated "people dont want faster internet", raised prices and put in place bandwidth caps. Alice pays $0.30 per GB for going over her 100 GB bandwidth cap.

Alice's ISP has been getting worse after net neutrality was struck down by a secret backdoor international trade agreement, that even members of congress were not allowed to see or vote on before it was signed.

Alice's Youtube and Netflix videos are loading slower than ever before. Alice's ISP has started throttling Netflix, Youtube and Bitorrent while publicly denying it. 

Alice's ISP has begun tracking every website she visits, recording her personal information and selling it to the NSA and marketing companies. Alice's ISP has been stealing revenue from the websites Alice visits, by replacing the website's ads with its own advertisements. Alice's ISP is starting to blacklist websites it doesnt like.

Alice hears about Skycoin, finds another Skycoin user with an office in the colocation center. Alice pays $1500 and installs 1.4 Gb/second Ubiquity airFiber hardware to bridge the distance between her and the fiber backbone.  Alice's connection acts as the backhaul for her neighborhood's local Skycoin mesh.

Alice cancels her internet service.

User Story: Internet Kill Switch
=================================

Todo