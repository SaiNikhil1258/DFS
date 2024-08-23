package main

import (
	"bytes"
	"log"
	"time"

	"guthub.com/sainikhil1258/DFS/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcptransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		// OnPeer: ,
	}
	tcpTransport := p2p.NewTCPTransport(tcptransportOpts)
	fileServerOpts := FileServerOpts{
		EncKey:            newEncryptionKey(),
		StorgaeRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		BootstrapNodes:    nodes,
		Transport:         tcpTransport,
	}
	s := NewFileServer(fileServerOpts)
	tcpTransport.OnPeer = s.OnPeeer
	return s
}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")
	go func() {
		log.Fatal(s1.Start())
	}()
	time.Sleep(1 * time.Second)
	go s2.Start()
	time.Sleep(1 * time.Second)

	data := bytes.NewReader([]byte("My big Data file here!"))
	s2.Store("CoolPicture.jpg", data)
	time.Sleep(5 * time.Millisecond)

	// r, err := s2.Get("CoolPicture.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// b, err := io.ReadAll(r)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(b))
	// select {}
}
