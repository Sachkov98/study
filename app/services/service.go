package services

import (
	"fmt"
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
	ConnectToDb() error
	InsertOrders([]order.Order) error
}

func (s service) Start() {

	for range time.Tick(time.Second * 60) {

		orders, err := s.gateway.GetOrders()
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = s.repository.InsertOrders(orders)
		if err != nil {
			fmt.Println(err)
		}
	}
}
