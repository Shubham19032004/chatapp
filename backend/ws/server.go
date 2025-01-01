package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	// Setup request size
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// Manager manages WebSocket connections
type Manager struct {
	clients ClientList
	sync.RWMutex
}

// NewManager creates a new WebSocket manager
func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

// ServerWS handles WebSocket connections (exported for external use)
func (m *Manager) ServerWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	client := NewClient(conn, m)
	m.addClinet(client)
	go client.readMessages()
	// defer conn.Close() // Ensure the connection is closed when done
	// log.Println("Connection closed")
}

func (m *Manager) addClinet(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true

}
func (m *Manager) removeClinet(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}

}
