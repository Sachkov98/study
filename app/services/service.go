package services

import (
	"study/app/domain/order"
	"time"
)

type service struct {
	gateway    OrdersGateway
	repository OrdersRepository
}

func New(gateway OrdersGateway, repository OrdersRepository) *service {
	serv := service{gateway, repository}
	return &serv
}

type OrdersGateway interface {
	GetOrders() ([]order.Order, error)
}

type OrdersRepository interface {
	InsertOrders([]order.Order) error
}

func (s service) GetOrdersInsertOrders() error {
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

func (s service) Start() {

	for range time.Tick(time.Second * 60) {
		s.GetOrdersInsertOrders()
	}

}
