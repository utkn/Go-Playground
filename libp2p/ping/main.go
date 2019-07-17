package main

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/multiformats/go-multiaddr"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	// Create a node that listens 2000 tcp port on 127.0.0.1.
	node, err := libp2p.New(ctx,
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"),
		libp2p.Ping(false),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Close the node
	defer func() {
		if err := node.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	pingService := &ping.PingService{Host: node}
	node.SetStreamHandler(ping.ID, pingService.PingHandler)

	fmt.Println("Listen addresses:", node.Addrs())

	peerInfo := &peer.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}

	addrs, err := peer.AddrInfoToP2pAddrs(peerInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("node address: ", addrs[0])

	if len(os.Args) > 1 {
		addr, _ := multiaddr.NewMultiaddr(os.Args[1])
		pr, _ := peer.AddrInfoFromP2pAddr(addr)
		err = node.Connect(ctx, *pr)
		ch := pingService.Ping(ctx, pr.ID)
		for i := 0; i < 10; i++ {
			res := <-ch
			fmt.Println("got ping response. RTT: ", res.RTT)
		}
	} else {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		fmt.Println("Received close signal from OS...")
	}
}
