package services

import (
	"fmt"
	"github.com/Sachkov98/study/app/adapters/primary/http-adapter/controller"
	"github.com/Sachkov98/study/app/domain/order"
	"time"
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
	GetOrdersByIds(controller.OrdersIds) ([]order.Order, error)
}

func (s Service) GetOrdersInsertOrders() error {
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
		err := s.GetOrdersInsertOrders()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func (s Service) GetOrders(ordersids controller.OrdersIds) ([]order.Order, error) {
	ordersIds, err := s.repository.GetOrdersByIds(ordersids)
	if err != nil {
		return []order.Order{}, err
	}
	return ordersIds, nil
}
