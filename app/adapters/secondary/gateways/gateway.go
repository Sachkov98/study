package gateways

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Sachkov98/study/app/domain/order"
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
	const url = "http://localhost:8081"

	context := context.Background()

	req, err := http.NewRequestWithContext(context, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, _ := g.client.Do(req)

	var dto DTO

	err = json.NewDecoder(response.Body).Decode(&dto)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err == nil {
			err = response.Body.Close()
			return
		}
	}()

	return dto.Orders, nil
}
