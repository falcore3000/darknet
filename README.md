

Darknet Plan
============

The Skycoin darknet is designed for a world where net neutrality has failed. A world where monoplistic government granted monopolies have become the gate keepers to the internet.

The Skycoin darknet emphasizes security, speed and privacy and is designed specificly for deploying open access wifi mesh networks and community ISPs.

The Skycoin routing layer is a technological response to ACTA, PIPA, SOPA and the Comcast Time Warner Merger. The proticol is designed to bridge the last mile between fiber and the home and eliminate reliance upon monopolistic ISPs.

Network deployment will begin by July.

Implementation Details:
- prototype in Golang
- extremely simple. less than 2000 lines of code
- minimal number of dependencies

Design Goals:
- Open Access Wifi mesh networks
- resistant to latency, high packet loss and low reliability connections
- coin incentives for provisioning bandwidth, storage and backhaul
- Designed to bridge last mile between the network backbone and home
- runs on Rasberry Pi and Ubibuity Hardware
- "zeroconf". Plug in and runs, no configuration
- difficult to detect and throttle

Technical Aspects:
- uses pubkey hashes as network addresses
- instant, low overhead, distributed bandwidth micropayments using off blockchain transactions
- link layer, does not define routing
- store and forward
- compatibility bridge with IPv6 networks
- Link Aggregation (ability to aggregate bandwidth from multiple connections)

Security:
- Link layer encryption between nodes
- End to End Encryption
- Post Public Key encryption primitives

User Stories: Link Aggregation
===============================

Bob has a 2 Mb/s internet connection and cannot watch Youtube videos. Bob and his five neighbors with 2 Mb/s connections install Skycoin nodes


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

Alice's ISP has been getting worse after net neutrality was struck down by a secret backdoor international trade agreement, which supercedes the constition and that even members of congress were not allowed to see or vote on before it was signed.

Alice's Youtube and Netflix videos are loading slower than ever before. Alice's ISP has started throttling Netflix, Youtube and Bitorrent while publicly denying it. 

Alice's ISP has begun tracking every website she visits, recording her personal information and selling it to the NSA and marketing companies. Alice's ISP has been stealing revenue from the websites Alice visits, by replacing the website's ads with its own advertisements. Alice's ISP is starting to blacklist websites it doesnt like.

Alice hears about Skycoin, finds another Skycoin user with an office in the colocation center. Alice pays $1500 and installs a 1.4 Gb/second, 24 GHz Ubiquiti airFiber to bridge the distance between her and the fiber backbone.  Alice's connection acts as the backhaul for her neighborhood's local Skycoin mesh.

Alice cancels her internet service.