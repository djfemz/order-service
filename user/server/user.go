package server

import (
	"context"

	"github.com/djfemz/user-service/db"
	"github.com/djfemz/user-service/proto/protos/user"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	logger         *logrus.Logger
	userRepository db.UserRepository
	user.UnimplementedUserServer
}

func NewUserService(logger *logrus.Logger, userRepo db.UserRepository) user.UserServer {
	return &UserService{logger: logger, userRepository: userRepo}
}

func (userService *UserService) GetUser(ctx context.Context, userRequest *user.UserRequest) (*user.UserResponse, error) {
	userService.logger.Info("In Get user for request:: ", userRequest)
	foundUser, err := userService.userRepository.GetUserById(uint64(userRequest.GetId()))
	if err != nil {
		userService.logger.Error("Error finding user with id:: ", userRequest.Id)
		return nil, err
	}
	return &user.UserResponse{Id: int32(foundUser.Id), Username: foundUser.Username, CreatedAt: foundUser.CreatedAt}, nil
}
