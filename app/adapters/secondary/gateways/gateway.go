package gateways

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type gateway struct{}

func New() *gateway {
	prov := gateway{}
	return &prov
}

type Order struct {
	OrderId     int64  `json:"order_id"`
	Status      string `json:"status"`
	StoreId     int64  `json:"store_id"`
	DateCreated string `json:"date_created"`
}

type DTO struct {
	Orders []Order `json:"content"`
}

func (g gateway) GetOrders() ([]Order, error) {

	var dto DTO

	resp, err := http.Get("http://localhost:8081")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &dto)
	if err != nil {
		log.Fatal(err)
	}
	return dto.Orders, err
}
