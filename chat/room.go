package chat

import (
	"log"
)

type room struct {
	forwardCh chan []byte
	joinCh    chan *client
	leaveCh   chan *client
	// TODO: map[int]*clientのほうが良い?
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {

		case client := <-r.joinCh:
			r.join(client)

		case client := <-r.leaveCh:
			r.leave(client)

		case message := <-r.forwardCh:
			log.Println("受け取ったメッセージ:", message)

			// すべてのクライアントにメッセージを送信.
			for client := range r.clients {
				select {
				// TODO: メソッド化
				case client.sendCh <- message:
					// send message
				default:
					// fail
					r.leave(client)
				}
			}

		}
	}
}

func (r *room) join(c *client) {
	log.Println("room join")
	r.clients[c] = true
}

func (r *room) leave(c *client) {
	log.Println("room leave")
	delete(r.clients, c)
	close(c.sendCh)
}
