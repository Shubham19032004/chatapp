package main

import (
	"database/sql"
	"log"

	"github.com/shubham19032004/chatapp/utils"
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
	store := db.NewStore(conn)
	runGrpcServer(config, store)

}
func runGrpcServer(config utils.Config,store db.Store){
	
}