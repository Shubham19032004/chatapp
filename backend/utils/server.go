package utils

import (
	"log"
	"net/http"

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
}

// NewManager creates a new WebSocket manager
func NewManager() *Manager {
	return &Manager{}
}

// ServerWS handles WebSocket connections (exported for external use)
func (m *Manager) ServerWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close() // Ensure the connection is closed when done
	log.Println("Connection closed")
}
