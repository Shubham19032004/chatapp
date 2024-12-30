package main

import (
	"database/sql"
	"log"
	"net"
	"github.com/shubham19032004/chatapp/utils"
	"github.com/shubham19032004/chatapp/gapi"
	"github.com/shubham19032004/chatapp/pb"
	db "github.com/shubham19032004/chatapp/db/sqlc"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func  main()  {
	config,err:=utils.LoadConfig(".")
	if err!=nil{
		log.Fatal("cannnot load config:",err)
	}
	// sql.open start a connection for db
	conn,err:=sql.Open(config.DBDriver,config.DBSource)
	if err!=nil{
		log.Fatal("connot connect to db:",err)
	}
	// provide method to execute database query
	store := db.NewStore(conn)
	runGrpcServer(config, store)

}

func runGrpcServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChatAppServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	log.Printf("start gPRC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc:", err)
	}
}