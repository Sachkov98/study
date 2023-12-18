package services

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {

	gatewayMock := NewGatewayMock()
	repositoryMock := NewRepositoryMock()
	service := New(gatewayMock, repositoryMock)

	err := service.GetOrdersInsertOrders()
	if err != nil {
		t.Fail()
	} else {
		fmt.Println("Ok!")
	}

	service.Start()
}
