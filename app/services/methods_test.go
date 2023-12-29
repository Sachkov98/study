package services

import (
	"testing"
)

func TestStart(t *testing.T) {

	gatewayMock := NewGatewayMock()
	repositoryMock := NewRepositoryMock()
	service := New(gatewayMock, repositoryMock)

	err := service.GetOrdersInsertOrders()
	if err != nil {
		t.Fail()
	}
}
