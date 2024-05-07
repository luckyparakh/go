package main

import (
	"dfs/p2p"
	"fmt"
	"log"
)

func onPeer(p p2p.Peer) error {
	p.Close()
	return nil
}
func main() {
	options := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		HandShaker:    p2p.NOPHandShake,
		Decoder:       p2p.DefaultDecode{},
		OnPeer:        onPeer,
	}
	tr := p2p.NewTCPTransport(options)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}

//toStart 2:18:02
