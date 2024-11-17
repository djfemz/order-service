package server

import (
	"context"

	"github.com/djfemz/order-service/proto/protos/order"
	

	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	logger *logrus.Logger
	userClient user.UserClient
}

func NewOrderService(logger *logrus.Logger, userClient user.UserClient) *OrderService {
	return &OrderService{
		logger: logger,
		userClient: userClient,
	}
}

func (orderService *OrderService) CreateOrder(ctx context.Context, orderRequest *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	orderService.logger.Info("In Create Order with request:: ", orderRequest)
	return nil, nil
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
