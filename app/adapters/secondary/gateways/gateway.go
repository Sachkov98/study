package gateways

import (
	"encoding/json"
	"github.com/Sachkov98/study/app/domain/order"
	"net/http"
)

type Gateway struct{}

func New() *Gateway {
	gateway := Gateway{}
	return &gateway
}

type DTO struct {
	Orders []order.Order `json:"content"`
}

func (g Gateway) GetOrders() ([]order.Order, error) {
	var dto DTO
	client := &http.Client{}
	resp, err := client.Get("http://localhost:8081")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&dto)
	if err != nil {
		return nil, err
	}
	return dto.Orders, nil
}
