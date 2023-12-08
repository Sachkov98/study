package services

import (
	"study/app/adapters/secondary/gateways"
	"study/app/adapters/secondary/repositories"
	"time"
)

type service struct {
}

func New(a, b interface{}) *service {
	serv := service{}
	return &serv
}

type GatewayInterface interface {
	GetOrders()
}

type RepositoryInterface interface {
	GetOrders()
	ConnectToDb()
	InsertOrdersToDb()
}

func (s service) Start() {
	for {
		ord, _ := gateways.New().GetOrders()
		db, _ := repositories.New().ConnectToDb()
		repositories.New().InsertOrdersToDb(ord, db)
		time.Sleep(60 * time.Second)
	}

}
