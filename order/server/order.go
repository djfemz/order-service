package server

import (
	"context"
	"sync"

	"github.com/djfemz/order-service/db"
	"github.com/djfemz/order-service/models"
	"github.com/djfemz/order-service/proto/protos/order"

	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	order.UnimplementedOrderServer
	logger     *logrus.Logger
	userClient user.UserClient
	orderRepo  db.OrderRepository
}

func NewOrderService(logger *logrus.Logger, userClient user.UserClient, orderRepo db.OrderRepository) order.OrderServer {
	return &OrderService{
		logger:     logger,
		userClient: userClient,
		orderRepo:  orderRepo,
	}
}

func (orderService *OrderService) CreateOrder(ctx context.Context, orderRequest *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	orderService.logger.Info("In Create Order with request:: ", orderRequest)
	var wg sync.WaitGroup
	var userDetails *user.UserResponse
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		userDetails, err = orderService.userClient.GetUser(ctx, &user.UserRequest{Id: orderRequest.UserId})
		orderService.logger.Info("User found:: ", userDetails)
	}()
	wg.Wait()
	if err != nil {
		orderService.logger.Error("Error fetching user :: ", err)
	}

	newOrder := models.NewOrder(orderRequest.Item, float64(orderRequest.Price))
	newOrder, err = orderService.orderRepo.Save(newOrder)
	if err != nil {
		return nil, err
	}
	orderService.logger.Info("user: ", userDetails)
	resp:=mapOrderToOrderResponse(newOrder, userDetails)
	orderService.logger.Info("resp:: ", resp)
	return resp, nil
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

func mapOrderToOrderResponse(createdOrder *models.Order, userDetails *user.UserResponse) *order.CreateOrderResponse {
	return &order.CreateOrderResponse{
		OrderId:   int32(createdOrder.Id),
		Item:      createdOrder.Item,
		Price:     float32(createdOrder.Price),
		CreatedAt: createdOrder.CreatedAt,
		CreatedBy: &order.GetUserResponse{
			Id: userDetails.Id,
			Username: userDetails.Username,
		},
		Status:    order.Status_COMPLETE,
	}
}
