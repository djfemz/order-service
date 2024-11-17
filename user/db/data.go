package db

import (
	"errors"
	"time"

	"github.com/djfemz/user-service/models"
	"github.com/sirupsen/logrus"
)

var users []*models.User = []*models.User{
	{Id: 1, Username: "john@email.com", CreatedAt: time.Now().String()}, 
	{Id: 2, Username: "jane@email.com", CreatedAt: time.Now().String()}, 
	{Id: 3, Username: "jon@email.com", CreatedAt: time.Now().String()}, 
	{Id: 4, Username: "johnny@email.com", CreatedAt: time.Now().String()},
}

type UserRepository interface{
	GetUserById(id uint64) (*models.User, error)
}

type UserRepositoryImpl struct {
	*logrus.Logger
}

func NewUserRepository(logger *logrus.Logger) UserRepository {
	return &UserRepositoryImpl{logger}
}

func (userRepo *UserRepositoryImpl) GetUserById(id uint64) (*models.User, error) {
	user:=findUser(id)
	if user!=nil{
		return user, nil
	}
	return nil, errors.New("failed to find user with id")
}

func findUser(id uint64) *models.User{
	for _, user := range users {
		if user.Id == id{
			return user
		}
	}
	return nil
}