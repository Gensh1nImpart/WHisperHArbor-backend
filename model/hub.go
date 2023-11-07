package model

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan *Message
	Register   chan *Client
	Unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *Message),
		Register:   make(chan *Client),
		Clients:    make(map[*Client]bool),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) getMessage(name string, Type int) {
	if Type == 1 {
		message := &Message{Msg: []byte(name + "上线了，当前在线人数: " + strconv.Itoa(len(h.Clients))), User: []byte("System"), Type: 0}
		go func() {
			h.mu.Lock()
			defer h.mu.Unlock()
			h.Broadcast <- message
		}()
	} else {
		message := &Message{Msg: []byte(name + "走了，当前在线人数: " + strconv.Itoa(len(h.Clients))), User: []byte("System"), Type: 0}
		go func() {
			h.mu.Lock()
			defer h.mu.Unlock()
			h.Broadcast <- message
		}()
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			//go func() {
			//	h.getMessage(string(client.User), 1)
			//}()
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				//go func() {
				//	h.getMessage(string(client.User), 0)
				//}()
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- *message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
			//go func() {
			//	h.HandleMessage(*message)
			//}()
		}
	}
}

var (
	AI_names = []string{"赵老师", "钱老师", "来大人", "孙老师", "小杰子", "Mslxl", "bt"}
)

func (h *Hub) HandleMessage(msg Message) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if strings.Contains(string(msg.Msg), "?") || strings.Contains(string(msg.Msg), "？") || strings.Contains(string(msg.Msg), "吗") {
		msg.Msg = []byte(strings.ReplaceAll(string(msg.Msg), "?", "!"))
		msg.Msg = []byte(strings.ReplaceAll(string(msg.Msg), "？", "!"))
		msg.Msg = []byte(strings.ReplaceAll(string(msg.Msg), "吗", ""))
		message := &Message{Msg: msg.Msg, User: []byte(AI_names[rand.Intn(len(AI_names))]), Type: 0}
		go func() {
			h.Broadcast <- message
		}()
	}

}
