package main

import (
	"log"
	"time"

	"guthub.com/sainikhil1258/DFS/p2p"
)

func main() {
	tcptransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcptransportOpts)
	FileServerOpts := FileServerOpts{
		StorgaeRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}
	s := NewFileServer(FileServerOpts)
	go func() {
		time.Sleep(time.Second * 3)
		s.Stop()
	}()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
