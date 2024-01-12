package services

import (
	"github.com/Sachkov98/study/app/domain/order"
)

type RepositoryMock struct{}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

func (rep RepositoryMock) InsertOrders([]order.Order) error {
	return nil
}

func (rep RepositoryMock) GetOrdersByIds([]int) ([]order.Order, error) {
	return nil, nil
}
