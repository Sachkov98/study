package services

import (
	"github.com/Sachkov98/study/app/domain/order"
)

type GatewayMock struct{}

func NewGatewayMock() *GatewayMock {
	return &GatewayMock{}
}

func (g GatewayMock) GetOrders() ([]order.Order, error) {
	return nil, nil
}
