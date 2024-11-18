package main

import (
	"net"
	"os"

	"github.com/djfemz/user-service/db"
	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/djfemz/user-service/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err:=os.OpenFile("user_service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err==nil{
		logger.Out = file
	}
	listener, err:=net.Listen("tcp", ":9001")
	if err!=nil{
		logger.Error("Error listening on port:: ", 9001)
	}
	userServer:=grpc.NewServer()
	user.RegisterUserServer(userServer, server.NewUserService(logger, db.NewUserRepository(logger)))
	logger.Info("Starting User service...")
	if err:=userServer.Serve(listener);err!=nil{
		logger.Error("error starting user server:: ", err)
	}
}