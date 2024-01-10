package gateways

import (
	"encoding/json"
	"github.com/Sachkov98/study/app/domain/order"
	"net/http"
)

type Gateway struct {
	client http.Client
}

func New() *Gateway {
	gateway := Gateway{}
	return &gateway
}

type DTO struct {
	Orders []order.Order `json:"content"`
}

func (g Gateway) GetOrders() ([]order.Order, error) {

	response, err := g.client.Get("http://localhost:8081")
	if err != nil {
		return nil, err
	}

	var dto DTO
	err = json.NewDecoder(response.Body).Decode(&dto)
	if err != nil {
		return nil, err
	}
	return dto.Orders, nil
}
