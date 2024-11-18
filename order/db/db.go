package db

import (
	"time"

	"github.com/djfemz/order-service/models"
	"github.com/sirupsen/logrus"
)

var orders []*models.Order = make([]*models.Order, 0)

type OrderRepository interface{
	Save(order *models.Order) (*models.Order, error)
}

type OrderRepositoryImpl struct {
	logger *logrus.Logger
}

func NewOrderRepository(logger *logrus.Logger) OrderRepository {
	return &OrderRepositoryImpl{
		logger: logger,
	}
}

func (orderRepo *OrderRepositoryImpl) Save(order *models.Order) (*models.Order, error) {
	order.Id = uint64(len(orders) + 1)
	order.CreatedAt=time.Now().String()
	orders = append(orders, order)
	orderRepo.logger.Info("adding new order:: ", order)
	return order, nil
}