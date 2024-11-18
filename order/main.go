package main

import (
	"net"
	"os"

	"github.com/djfemz/order-service/db"
	"github.com/djfemz/order-service/proto/protos/order"
	"github.com/djfemz/order-service/server"
	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":9002"

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("order_service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	}
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Error listening on port:: ", port)
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
