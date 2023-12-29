package gateways

import (
	"encoding/json"
	"github.com/Sachkov98/study/app/domain/order"
	"io/ioutil"
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

	resp, err := http.Get("http://localhost:8081")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, err
	}
	return dto.Orders, nil
}
