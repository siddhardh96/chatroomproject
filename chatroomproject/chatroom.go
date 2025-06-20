package main

import (
	"sync"
)

type Client struct {
	ID      string
	MsgChan chan string
}

//chatroom struct
type ChatRoom struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan string
	mu         sync.RWMutex
}

//everytime new chatroom will be created when ever a new client register
func NewChatRoom() *ChatRoom {
	room := &ChatRoom{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan string),
	}

	go room.run()
	return room
}

//this will run to send and recieve messages
func (cr *ChatRoom) run() {
	for {
		select {
		case client := <-cr.register:
			cr.mu.Lock()
			cr.clients[client.ID] = client
			cr.mu.Unlock()

		case client := <-cr.unregister:
			cr.mu.Lock()
			if _, ok := cr.clients[client.ID]; ok {
				delete(cr.clients, client.ID)
				close(client.MsgChan)
			}
			cr.mu.Unlock()

		case message := <-cr.broadcast:
			cr.mu.RLock()
			for _, client := range cr.clients {
				select {
				case client.MsgChan <- message:
				default:
				}
			}
			cr.mu.RUnlock()
		}
	}
}

//a new client join the meet
func (cr *ChatRoom) Join(id string) *Client {
	client := &Client{
		ID:      id,
		MsgChan: make(chan string, 100),
	}
	cr.register <- client
	return client
}

//this is to leave the room
func (cr *ChatRoom) Leave(id string) {
	cr.mu.RLock()
	client, ok := cr.clients[id]
	cr.mu.RUnlock()
	if ok {
		cr.unregister <- client
	}
}

//This one for send message to client
func (cr *ChatRoom) SendMessage(senderID, message string) {
	formatted := "[" + senderID + "]: " + message
	cr.broadcast <- formatted
}

//this one to get a client
func (cr *ChatRoom) GetClient(id string) (*Client, bool) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	client, ok := cr.clients[id]
	return client, ok
}
