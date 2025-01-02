package ws

import (
	"errors"
	"fmt"
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

	handlers map[string]EventHandler
}

// NewManager creates a new WebSocket manager
func NewManager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()
	return m
}
func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessage
}

func SendMessage(event Event, c *Client) error {
	fmt.Println(event)
	return nil
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("there is no such event")
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
	go client.writeMessages()
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
