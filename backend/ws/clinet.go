package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}
func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClinet(c)
	}()
	for {

		messagetype, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error reading Message:%v", err)
			}
			break
		}
		log.Println(messagetype)
		log.Println(string(payload))

	}
}
