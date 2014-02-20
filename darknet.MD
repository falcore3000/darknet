

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