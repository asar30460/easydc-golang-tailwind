package server

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       int    `json:"id"`
	ServerID int `json:"server_id"`
	Username string `json:"user_name"`
}

type Message struct {
	Content   string `json:"content"`
	ServerID  int    `json:"server_id"`
	ChannelID int `json:"channel_id"`
	Username  string `json:"user_name"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:   string(m),
			ServerID:  c.ServerID,
			ChannelID: c.ID,
			Username:  c.Username,
		}

		hub.Broadcast <- msg
	}
}
