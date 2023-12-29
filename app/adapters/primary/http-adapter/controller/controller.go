package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Sachkov98/study/app/domain/order"
	"net/http"
)

type Controller struct {
	service OrdersService
}

func New(service OrdersService) *Controller {
	controller := Controller{service}
	return &controller
}

type OrdersService interface {
	GetOrders(OrdersIds) ([]order.Order, error)
}

type OrdersIds struct {
	OrdersIds []int `json:"orders_ids"`
}

type DTO struct {
	Orders []order.Order `json:"orders"`
}

func (ctr Controller) GetOrders(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var ordersIds OrdersIds
	var dto DTO

	err := decoder.Decode(&ordersIds)
	if err != nil {
		fmt.Println("Error occured while decoding the data: ", err)
	}

	dto.Orders, err = ctr.service.GetOrders(ordersIds)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := json.Marshal(dto)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(req)

	return
}
