package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan Event
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan Event),
	}
}
func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClinet(c)
	}()
	for {

		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error reading Message:%v", err)
			}
			break
		}
		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Println("error marshalling event:%v", err)
			break
		}
		if err:=c.manager.routeEvent(request,c);err!=nil{
			log.Println("error handeling message:",err)
		}

	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.manager.removeClinet(c)
	}()
	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed:", err)
				}
				return
			}
			data,err:=json.Marshal(message)
			if err!=nil{
				log.Println(err)
				return 
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println("failed to send message: ", err)
			}
		}
	}
}
