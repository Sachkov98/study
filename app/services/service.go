package services

import (
	"study/app/adapters/secondary/gateways"
	"study/app/adapters/secondary/repositories"
)

type service struct {
}

func New() interface{} {
	serv := service{}
	return &serv
}

type GatewayInterface interface {
	GetBody()
}

type RepositoryInterface interface {
	DbInit()
	InsertListOrdersToDb()
}

func (s service) start() {
	gateways.New().GetBody()
	repositories.New().DbInit()
	repositories.New().InsertListOrdersToDb()

}
