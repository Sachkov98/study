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
	const url = "http://orders_api:8080"

	context := context.Background()

	req, err := http.NewRequestWithContext(context, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := response.Body.Close()
		if closeErr != nil {
			if err == nil {
				err = closeErr
			}
		}
	}()

	var dto DTO

	err = json.NewDecoder(response.Body).Decode(&dto)
	if err != nil {
		return nil, err
	}

	return dto.Orders, nil
}
