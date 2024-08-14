package main

import (
	"log"

	"guthub.com/sainikhil1258/DFS/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Hello raa poookaaa")
	select {}
}
