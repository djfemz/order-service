package server

import (
	"context"

	"github.com/djfemz/grpc-store/protos/user"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	logger *logrus.Logger
}



func (userService *UserService) GetUser(ctx context.Context, userRequest *user.UserRequest) (*user.UserResponse, error)  {
	userService.logger.Info("In Get user for request: ", userRequest)
	
	return nil, nil
}