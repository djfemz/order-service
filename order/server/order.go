package server

import (
	"context"
	"sync"

	apperrors "github.com/djfemz/order-service/appErrors"
	"github.com/djfemz/order-service/db"
	"github.com/djfemz/order-service/models"
	"github.com/djfemz/order-service/proto/protos/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	userDetailsCh := make(chan *user.UserResponse)
	errorCh := make(chan error)
	wg.Add(1)
	go func() {
		defer wg.Done()
		userDetails, err = orderService.userClient.GetUser(ctx, &user.UserRequest{Id: orderRequest.UserId})
		if err != nil {
			orderService.logger.Error("Error fetching user :: ", err)
			errorCh<-err
		}
		userDetailsCh<-userDetails
	}()
	go func() {
		wg.Wait()
		close(userDetailsCh)
		close(errorCh)
	}()
	select {
	case userDetails := <-userDetailsCh:
		orderService.logger.Info("user found with details:: ", userDetails)
	case err := <-errorCh:
		err = status.Newf(codes.InvalidArgument, apperrors.INVALID_USER_ID_ERROR_MESSAGE, orderRequest.GetUserId(), err.Error()).Err()
		return nil, err
	}

	newOrder, err := orderService.createOrder(orderRequest, err)
	if err!=nil {
		err = status.Newf(codes.InvalidArgument, apperrors.ORDER_CREATION_FAILED_MESSAGE, err.Error()).Err()
		return &order.CreateOrderResponse{
			Status: order.Status_FAILED,
		}, err
	}
	return mapOrderToOrderResponse(newOrder, userDetails), nil
}

func (orderService *OrderService) createOrder(orderRequest *order.CreateOrderRequest, err error) (*models.Order, error) {
	newOrder := models.NewOrder(orderRequest.Item, float64(orderRequest.Price))
	newOrder, err = orderService.orderRepo.Save(newOrder)
	if err != nil {
		orderService.logger.Error("Error creating order for request:: ", orderRequest)
		errorStatus := status.Newf(codes.InvalidArgument, apperrors.ORDER_CREATION_FAILED_MESSAGE, err.Error())
		errorStatus, _ = errorStatus.WithDetails(orderRequest)
		return nil, errorStatus.Err()
	}
	return newOrder, nil
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
	orderResponse := &order.CreateOrderResponse{
		OrderId:   int32(createdOrder.Id),
		Item:      createdOrder.Item,
		Price:     float32(createdOrder.Price),
		CreatedAt: createdOrder.CreatedAt,
		
		Status:    order.Status_COMPLETE,
	}

	if userDetails != nil{
		orderResponse.CreatedBy= &order.GetUserResponse{
			Id: userDetails.Id,
			Username: userDetails.Username,
		}
	}
	return orderResponse
}
