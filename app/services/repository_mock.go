package services

import (
	_ "github.com/lib/pq"
	"study/app/domain/order"
)

type RepositoryMock struct {
}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

func (rep RepositoryMock) InsertOrders(orders []order.Order) error {

	return nil
}
