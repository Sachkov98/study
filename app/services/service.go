package services

import (
	"study/app/adapters/secondary/providers"
	"study/app/adapters/secondary/repositories"
	"time"
)

type service struct {
}

func New(I) *service {
	serv := service{}
	return &serv
}

type I interface {
	GetBody()
	InsertTable()
}

var Intf I

func (s service) GetNDrop() {
	for {
		providers.New().GetBody()
		repositories.New().InsertTable()
		time.Sleep(60 * time.Second)
	}

}
