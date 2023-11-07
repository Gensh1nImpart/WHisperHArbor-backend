package model

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"time"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan Message
	User []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.Broadcast <- &Message{Msg: message, User: c.User, Type: 1}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			message.Msg = []byte(strconv.QuoteToASCII(string(message.Msg)))
			message.User = []byte(strconv.QuoteToASCII(string(message.User)))
			w, err := c.Conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			msg, err := json.Marshal(message)
			if err != nil {
				log.Println("Error in marshaling message:", err)
			} else {
				w.Write(msg)
			}
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				message = <-c.Send
				message.Msg = []byte(strconv.QuoteToASCII(string(message.Msg)))
				message.User = []byte(strconv.QuoteToASCII(string(message.User)))
				msg, err = json.Marshal(message)
				if err != nil {
					log.Println("Error in marshaling message:", err)
				} else {
					w.Write(msg)
				}
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
