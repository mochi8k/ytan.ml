package chat

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type room struct {
	forwardCh chan []byte
	joinCh    chan *client
	leaveCh   chan *client
	// TODO: map[int]*clientのほうが良い?
	clients map[*client]bool
}

func newRoom() *room {
	return &room{
		forwardCh: make(chan []byte),
		joinCh:    make(chan *client),
		leaveCh:   make(chan *client),
		clients:   make(map[*client]bool),
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// WebSocketコネクションを取得
	socket, err := upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Fatal("ServeHttp:", err)
		return
	}

	client := &client{
		socket: socket,
		sendCh: make(chan []byte, messageBufferSize),
		room:   r,
	}

	r.joinCh <- client

	defer func() {
		r.leaveCh <- client
	}()

	go client.write()

	client.read()
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
