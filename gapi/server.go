package gapi

import (
	"fmt"
	db "github.com/HoangLong2502/simple_bank/db/sqlc"
	"github.com/HoangLong2502/simple_bank/pb"
	"github.com/HoangLong2502/simple_bank/token"
	"github.com/HoangLong2502/simple_bank/utils"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     utils.Config
	store      *db.Store
	tokenMaker token.Maker
}

func NewServer(config utils.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
