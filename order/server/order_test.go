package server

import (
	"context"
	"errors"
	"testing"

	"github.com/djfemz/order-service/db"
	"github.com/djfemz/order-service/proto/protos/order"
	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockUserClient struct {
	user.UnimplementedUserServer
	mock.Mock
}

func (m *MockUserClient) GetUser(ctx context.Context, req *user.UserRequest, opts...grpc.CallOption) (*user.UserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*user.UserResponse), args.Error(1)
}

func TestCreateOrder_Success(t *testing.T) {
	mockUserClient := &MockUserClient{}
	mockUserClient.On("GetUser", mock.Anything, &user.UserRequest{Id: 1}).Return(&user.UserResponse{
		Id: 1,
		Username: "test user",
		CreatedAt:  "2024-11-11",
	}, nil)
	logger:=logrus.New()
	orderService := &OrderService{logger:logger, userClient: mockUserClient, orderRepo: db.NewOrderRepository(logger)}
	resp, err := orderService.CreateOrder(context.TODO(), &order.CreateOrderRequest{
		UserId:  1,
		Item: "Milk",
		Price: 50.00,
	})

	assert.NoError(t, err)
	assert.Equal(t, int32(1), resp.CreatedBy.Id)
	assert.Equal(t, order.Status_COMPLETE, resp.Status)
	mockUserClient.AssertExpectations(t)
}

func TestCreateOrder_Invalid_UserId_Failure(t *testing.T) {
	mockUserClient := &MockUserClient{}
	mockUserClient.On("GetUser", mock.Anything, &user.UserRequest{Id: 12}).Return(&user.UserResponse{}, errors.New("user with id not found"))
	logger:=logrus.New()
	orderService := &OrderService{logger:logger, userClient: mockUserClient, orderRepo: db.NewOrderRepository(logger)}
	resp, err := orderService.CreateOrder(context.TODO(), &order.CreateOrderRequest{
		UserId:  12,
		Item: "Milk",
		Price: 50.00,
	})

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockUserClient.AssertExpectations(t)
}