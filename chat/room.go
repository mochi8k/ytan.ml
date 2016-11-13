package main

type room struct {
	forwardCh chan []byte
	joinCh    chan *client
	leaveCh   chan *client
	clients   map[*client]bool
}
