package server

import (
	"context"

	"github.com/djfemz/grpc-store/protos/order"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	logger *logrus.Logger
	userService *UserService
}

func NewOrderService(logger *logrus.Logger, userService *UserService) *OrderService{
	return &OrderService{
		logger: logger,
		userService: userService,
	}
}

func (orderService *OrderService) CreateOrder(ctx context.Context, orderRequest *order.CreateOrderRequest) (*order.CreateOrderResponse, error)  {
	orderService.logger.Info("In Create Order with request:: ", orderRequest)
	return nil, nil
}

func (orderService *OrderService) GetUser(ctx context.Context, getOrderRequest *order.GetUserRequest) (*order.UserResponse, error){
	orderService.logger.Info("In Get User with request:: ", getOrderRequest)
	
	return nil, nil
}


