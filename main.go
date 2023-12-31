package main

import (
	"database/sql"
	"github.com/HoangLong2502/simple_bank/gapi"
	"github.com/HoangLong2502/simple_bank/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/HoangLong2502/simple_bank/api"
	db "github.com/HoangLong2502/simple_bank/db/sqlc"
	"github.com/HoangLong2502/simple_bank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)
}

func runGrpcServer(config utils.Config, store *db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func runGinServer(config utils.Config, store *db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
