package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/multiformats/go-multiaddr"
	"log"
	"os"
)

// This function is called whenever a node wants to connect to this node using the
// foo protocol.
func fooStreamHandler(s network.Stream) {
	fmt.Println("Receiving stream!!")
	reader := bufio.NewReader(s)
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Could not get the message: %v", err)
		return
	}
	fmt.Printf("Got message: %s", msg)
	// Close this end of the stream.
	err = s.Close()
	if err != nil {
		fmt.Printf("Could not close the stream: %v", err)
		return
	}
}

func main() {
	ctx := context.Background()

	if len(os.Args) < 2 {
		log.Fatalln("Please provide a port to listen!")
	}

	listenAddr := fmt.Sprintf("/ip4/127.0.0.1/tcp/%s", os.Args[1])

	// Create a host and listen
	host, err := libp2p.New(ctx,
		libp2p.ListenAddrStrings(listenAddr))

	if err != nil {
		log.Fatalf("Could not create a host: %v", err)
	}

	defer host.Close()

	hostInfo := &peer.AddrInfo{
		Addrs: host.Addrs(),
		ID:    host.ID(),
	}

	// Convert hostInfo to an Multiaddr array
	addrs, _ := peer.AddrInfoToP2pAddrs(hostInfo)

	// Print out the multiaddress of this node
	fmt.Println(addrs[0])

	host.SetStreamHandler("/foo/1.0.0", fooStreamHandler)

	for {
		fmt.Println("Enter address to connect:")
		var targetAddress string
		_, err := fmt.Scanln(&targetAddress)
		if err != nil {
			fmt.Printf("Could not parse the address: %v\n", err)
			continue
		}
		// Generate a multiaddress from the input.
		multiAddr, err := multiaddr.NewMultiaddr(targetAddress)
		if err != nil {
			fmt.Printf("Could not parse the address: %v\n", err)
			continue
		}

		// Generate address info from the multi address.
		targetInfo, err := peer.AddrInfoFromP2pAddr(multiAddr)
		if err != nil {
			fmt.Printf("Could not parse the address: %v\n", err)
			continue
		}

		// Add the target info to this host's peer store.
		host.Peerstore().AddAddrs(targetInfo.ID, targetInfo.Addrs, peerstore.PermanentAddrTTL)

		// Generate a stream with the target.
		s, err := host.NewStream(ctx, targetInfo.ID, "/foo/1.0.0")
		if err != nil {
			fmt.Printf("Could not generate the stream: %v\n", err)
			continue
		}

		// Write a hello message to the stream.
		writer := bufio.NewWriter(s)
		_, err = writer.WriteString(fmt.Sprintf("hello from %v\n", host.ID()))
		if err != nil {
			fmt.Printf("Could not send the message: %v\n", err)
			continue
		}
		writer.Flush()
		// Close the stream
		//s.Close()
	}
}
