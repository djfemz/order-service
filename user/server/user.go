package server

import (
	"context"
	"time"

	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	logger *logrus.Logger
	user.UnimplementedUserServer
}

func NewUserService(logger *logrus.Logger) *UserService {
	return &UserService{logger: logger}
}

func (userService *UserService) GetUser(ctx context.Context, userRequest *user.UserRequest) (*user.UserResponse, error) {
	userService.logger.Info("In Get user for request: ", userRequest)

	return &user.UserResponse{Id: 1, Username: "john@email.com", CreatedAt: time.Now().GoString()}, nil
}
