package main

import (
	"net"

	"github.com/djfemz/user-service/db"
	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/djfemz/user-service/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logger:=logrus.New()
	listener, err:=net.Listen("tcp", ":9001")
	if err!=nil{
		logger.Error("Error listening on port:: ", 9001)
	}
	userServer:=grpc.NewServer()
	user.RegisterUserServer(userServer, server.NewUserService(logger, db.NewUserRepository(logger)))
	userServer.Serve(listener)
}