package gapi
import (
	// "context"
	// "fmt"
	db "github.com/shubham19032004/chatapp/db/sqlc"
	"github.com/shubham19032004/chatapp/pb"
	// "github.com/shubham19032004/chatapp/token"
	"github.com/shubham19032004/chatapp/utils"
    "google.golang.org/grpc"

)

type Server struct{
    pb.UnimplementedChatAppServer
    store      db.Store
	// tokenMaker token.Maker
	config     utils.Config
}

func NewServer(config utils.Config,store db.Store) (*Server,error){
    server := &Server{
		config:     config,
		// tokenMaker: tokenMaker,
	}
    return server, nil

}
func (s *Server) HelloServer(stream grpc.BidiStreamingServer[pb.Send, pb.Receive]) error {
    // Implement the method here to handle the stream
	return nil
}

