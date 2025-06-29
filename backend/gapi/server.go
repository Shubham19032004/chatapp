package gapi

import (

	
	"fmt"
	"io"
	"log"

	db "github.com/shubham19032004/chatapp/db/sqlc"
	"github.com/shubham19032004/chatapp/pb"

	// "github.com/shubham19032004/chatapp/token"
	"github.com/shubham19032004/chatapp/utils"
)

type Server struct {
	pb.UnimplementedChatAppServer
	store db.Store
	// tokenMaker token.Maker
	config utils.Config
}


func NewServer(config utils.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		// tokenMaker: tokenMaker,
	}
	return server, nil

}


func (s *Server) HelloServer(stream pb.ChatApp_HelloServerServer) error {
	log.Println("Bidirectional stream started")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("HelloServer: Client has closed the stream")
			return nil
		}
		if err != nil {
			log.Printf("HelloServer: Error receiving message: %v\n", err)
			log.Printf("HelloServer: Received message: %s\n", req.Message)

			// Send a response to the client
			response := &pb.Receive{
				Message: fmt.Sprintf("Server acknowledges: %s", req.Message),
			}
			if err := stream.Send(response); err != nil {
				log.Printf("HelloServer: Error sending response: %v\n", err)
				return err
			}
		}
	}

}
