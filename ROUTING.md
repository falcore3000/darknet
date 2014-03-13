Use-cases
=========

1.  I can get access to a fast line somewhere close to where I need it.  Wifi the last mile.  Make Mom&Pop single-hop ISP to finance the line and the equipment.

2.  I can get access to a REALLY fast line somewhere close to where I need it.  Wifi the last mile.  Mom&Pop ISP where I setup multiple hops myself or I get others to participate in the last mile in order to achieve the link from my line to my house/work.

3.  I have fastish internet at home and so do all my neighbors.  We all have the same ISP.  We can help each other out with peak loads by locating a computer close to our ISP with a fat pipe which can stream in at a high rate and then split it into multiple streams which all go down our respective connections to be re-assembled by my computer.

4.  I have fastish internet at home and so do all my neighbors.  We all have internet from the duopoloy or triopoly in our region.  We need a computer somewhere on the 'net which has good bandwidth to both (all) ISPs in order to perform channel unbonding and rebonding.

5.  I operate a network that blankets an area in wifi which is backed up by some kind of reliable backhaul.  I want to provide extra speed.  If I run darknet gateways on all my access points and a darknet aggregator somewhere near my POP people can take advantage of channel bonding to get extra speed.


Names
=====

darknet node - potentially anything which participates in the darknet
darknet client - a client which ostensibly does not have access to clearnet (though in some cases it might)
darknet gateway - something which bridges the darknet to clearnet
darknet aggregator - a server which sits on very fast connections in order to assist with one side of the channel bonding in the cases which require it

Design strategy
===============

We're not trying to build a mesh-net that could allow anyone to contact anyone always and forever.  We're trying to build something that makes the last mile much faster for people.  That means we can assume that every darknet terminates at one gateway.  And if we make that assumption every darknet turns into a tree.  That simplifies routing in an immense way.

There are two layers of darknet.

1.  "Link" layer.  This is a way of getting packets from the darknet client to the clearnet gateway and back to the darknet client.  Some networks don't need any more than this if they have very fast connections.

2.  "Bonding" layer.  This is a way of coordinating many packets to be routed from a a single darknet client to multiple darknet gateways and then on to a darknet aggregator.  Then back from the aggregator to multiple darknet gateways and on to a single darknet client.

If someone has a really fat pipe link layer is really all that's needed, it's a way to earn credit for backhaul.  Once it gets to their darknet gateway it's out on the clearnet and everything is OK.

<todo> Fix these names, as they're not right
