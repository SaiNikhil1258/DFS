package main

import (
	"fmt"
	"log"

	"guthub.com/sainikhil1258/DFS/p2p"
)

type FileServerOpts struct {
	// ListenAddr        string
	StorgaeRoot       string
	PathTransformFunc PathTransformFunc
	Transport         p2p.Transport
	// TCPTransportOpts  p2p.TCPTransportOpts
}

type FileServer struct {
	FileServerOpts
	store  *Store
	quitch chan struct{}
}

func NewFileServer(opts FileServerOpts) *FileServer {
	StoreOpts := StoreOpts{
		Root:              opts.StorgaeRoot,
		PathTransformFunc: opts.PathTransformFunc,
	}
	return &FileServer{
		FileServerOpts: opts,
		store:          NewStore(StoreOpts),
		quitch:         make(chan struct{}),
	}
}

func (s *FileServer) Stop() {
	close(s.quitch)
}

func (s *FileServer) loop() {
	defer func() {
		log.Println("file server stopped due to user quit action")

	}()
	for {
		select {
		case msg := <-s.Transport.Consume():
			fmt.Println(msg)
		case <-s.quitch:
			return
		}
	}
}

func (s *FileServer) Start() error {
	if err := s.Transport.ListenAndAccept(); err != nil {
		return err
	}
	s.loop()
	return nil
}
