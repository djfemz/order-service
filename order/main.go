package main

import (
	"net"

	"github.com/djfemz/order-service/db"
	"github.com/djfemz/order-service/proto/protos/order"
	"github.com/djfemz/order-service/server"
	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := logrus.New()
	listener, err := net.Listen("tcp", ":9002")
	if err != nil {
		logger.Error("Error listening on port:: ", 9002)
	}
	grpcClient, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("ERROR:: ", err)
	}
	userClient := user.NewUserClient(grpcClient)
	orderServer := grpc.NewServer()
	order.RegisterOrderServer(orderServer, server.NewOrderService(logger, userClient, db.NewOrderRepository(logger)))
	if err := orderServer.Serve(listener); err != nil {
		logger.Error("error starting server:: ", err)
	}
}
