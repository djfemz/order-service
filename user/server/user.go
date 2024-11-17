package server

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/djfemz/user-service/proto/protos/user"
)

type UserService struct {
	logger *logrus.Logger
}

func NewUserService(logger *logrus.Logger) *UserService {
	return &UserService{logger: logger}
}

func (userService *UserService) GetUser(ctx context.Context, userRequest *user.UserRequest) (*user.UserResponse, error) {
	userService.logger.Info("In Get user for request: ", userRequest)

	return nil, nil
}
