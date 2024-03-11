package websockets

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Manager struct {
	activeClients ActiveClientList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New websocket connection")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	conn.Close()
}

func (m *Manager) addClient(c *Client) {
	m.Lock()
	defer m.Unlock()
	m.activeClients[c] = true
}

func (m *Manager) removeClient(c *Client) {
	m.Lock()
	defer m.Unlock()
	delete(m.activeClients, c)
}