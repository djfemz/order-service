package main

import (
	"context"

	"github.com/djfemz/order-service/proto/protos/order"
	"github.com/djfemz/order-service/server"
	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := logrus.New()
	grpcClient, err:=grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil{
		logger.Error("ERROR:: ", err)
	}
	userClient:=user.NewUserClient(grpcClient)
	orderService:=server.NewOrderService(logger, userClient)
	getUserRequest:= &order.GetUserRequest{Id: 1}
	user, err:=orderService.GetUser(context.TODO(), getUserRequest)
	if err!= nil{
		logger.Error("ERROR:: ", err)
	}
	logger.Info("Found User:: ", user)
}