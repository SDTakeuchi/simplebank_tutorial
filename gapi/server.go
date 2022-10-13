package gapi

import (
	"fmt"
	db "simplebank/db/sqlc"
	"simplebank/pb"
	"simplebank/token"
	"simplebank/util"
)

type Server struct {
	pb.UnimplementedSimplebankServer
	config     util.Config
	store      db.Store // used to be "store  *db.Store" (with star) because it was defined as struct instaed of interface
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}