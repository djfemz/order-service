package server

import (
	"context"
	"time"

	"github.com/djfemz/order-service/db"
	"github.com/djfemz/order-service/models"
	"github.com/djfemz/order-service/proto/protos/order"

	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	order.UnimplementedOrderServer
	logger *logrus.Logger
	userClient user.UserClient
	orderRepo db.OrderRepository
}

func NewOrderService(logger *logrus.Logger, userClient user.UserClient, orderRepo db.OrderRepository) order.OrderServer {
	return &OrderService{
		logger: logger,
		userClient: userClient,
		orderRepo: orderRepo,
	}
}

func (orderService *OrderService) CreateOrder(ctx context.Context, orderRequest *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	orderService.logger.Info("In Create Order with request:: ", orderRequest)
	newOrder:=&models.Order{
		Item: orderRequest.Item,
		Price: float64(orderRequest.Price),
		CreatedAt: time.Now().String(),
	}
	newOrder, err:=orderService.orderRepo.Save(newOrder)
	if err!= nil{
		return nil, err
	}

	return &order.CreateOrderResponse{
		OrderId: int32(newOrder.Id),
		Item: newOrder.Item,
		Price: float32(newOrder.Price),
		CreatedAt: newOrder.CreatedAt,
	}, nil
}

func (orderService *OrderService) GetUser(ctx context.Context, getUserRequest *order.GetUserRequest) (*order.GetUserResponse, error) {
	orderService.logger.Info("In Get User with request:: ", getUserRequest)
	userRequest := &user.UserRequest{Id: getUserRequest.GetId()}
	resp, err := orderService.userClient.GetUser(context.TODO(), userRequest)
	if err != nil {
		orderService.logger.Error("Error fetching user from user service with ERROR:: ", err)
		return nil, err
	}
	return mapUserToUserResponse(resp), nil
}

func mapUserToUserResponse(user *user.UserResponse) *order.GetUserResponse {
	return &order.GetUserResponse{
		Id:        user.Id,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}
