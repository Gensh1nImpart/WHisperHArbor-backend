package model

import (
	"math/rand"
	"strconv"
	"strings"
)

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan *Message
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *Message),
		Register:   make(chan *Client),
		Clients:    make(map[*Client]bool),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) getMessage(name string, Type int) string {
	if Type == 1 {
		return name + "上线了，当前在线人数: " + strconv.Itoa(len(h.Clients))
	} else {
		return name + "走了，当前在线人数: " + strconv.Itoa(len(h.Clients))
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			message := &Message{Msg: []byte(h.getMessage(client.User.Account, 1)), User: []byte("System"), Type: 0}
			go func() {
				h.Broadcast <- message
			}()
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				message := &Message{Msg: []byte(h.getMessage(client.User.Account, 0)), User: []byte("System"), Type: 0}
				go func() {
					h.Broadcast <- message
				}()
				close(client.Send)
			}

		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- *message:
				default:
					close(client.Send)
					message := &Message{Msg: []byte(h.getMessage(client.User.Account, 0)), User: []byte("System"), Type: 0}
					go func() {
						h.Broadcast <- message
					}()
					delete(h.Clients, client)
				}
			}
			go func() {
				h.HandleMessage(*message)
			}()
		}
	}
}

var (
	AI_names = []string{"赵老师", "钱老师", "来大人", "孙老师", "小杰子", "Mslxl", "bt"}
)

func (h *Hub) HandleMessage(msg Message) {
	if strings.Contains(string(msg.Msg), "?") || strings.Contains(string(msg.Msg), "？") || strings.Contains(string(msg.Msg), "吗") {
		msg.Msg = []byte(strings.ReplaceAll(string(msg.Msg), "?", "!"))
		msg.Msg = []byte(strings.ReplaceAll(string(msg.Msg), "？", "!"))
		msg.Msg = []byte(strings.ReplaceAll(string(msg.Msg), "吗", ""))
	}
	h.Broadcast <- &Message{Msg: msg.Msg, User: []byte(AI_names[rand.Intn(len(AI_names))]), Type: 0}
}
