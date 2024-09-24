package main

import (
	"log"

	"github.com/Reneechang17/Distributed-File-System/p2p"
)

func main() {
	tcptransportOpts := p2p.TCPTransportOpts {
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
		// TODO: Implement OnPeer
	}
	tcpTransport := p2p.NewTCPTransport(tcptransportOpts)

	fileServerOpts := FileServerOpts {
		StorageRoot: "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport: tcpTransport,
	}

	s := NewFileServer(fileServerOpts)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}