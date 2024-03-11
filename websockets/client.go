package websockets

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id      uuid.UUID
	conn    *websocket.Conn
	manager *Manager
	send    chan []byte
}

type ActiveClientList map[*Client]bool

func NewClient(c *websocket.Conn, m *Manager) *Client {
	return &Client{
		id:      uuid.New(),
		conn:    c,
		manager: m,
		send:    make(chan []byte),
	}
}

func (c *Client) readMsg() {
	defer func() {
		c.manager.removeClient(c)
		c.conn.Close()
	}()
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Fatalln(err)
			}
		}
	}
}
