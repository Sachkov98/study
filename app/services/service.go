package services

import (
	"log"
	"time"

	"github.com/Sachkov98/study/app/domain/order"
)

type Service struct {
	gateway    OrdersGateway
	repository OrdersRepository
}

func New(gateway OrdersGateway, repository OrdersRepository) *Service {
	service := Service{gateway, repository}

	return &service
}

type OrdersGateway interface {
	GetOrders() ([]order.Order, error)
}

type OrdersRepository interface {
	InsertOrders([]order.Order) error
	GetOrdersByIds([]int) ([]order.Order, error)
}

func (s Service) getOrdersInsertOrders() error {
	orders, err := s.gateway.GetOrders()
	if err != nil {
		return err
	}

	err = s.repository.InsertOrders(orders)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Start() {
	for range time.Tick(time.Second * 60) {
		err := s.getOrdersInsertOrders()
		if err != nil {
			log.Println(err)

			continue
		}
	}
}

func (s Service) GetOrders(ordersIds []int) ([]order.Order, error) {
	orders, err := s.repository.GetOrdersByIds(ordersIds)
	if err != nil {
		return []order.Order{}, err
	}

	return orders, nil
}
