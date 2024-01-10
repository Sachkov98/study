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
	GetOrders([]int) ([]order.Order, error)
}

type ordersIdsDTO struct {
	OrdersIds []int `json:"orders_ids"`
}

type ordersDTO struct {
	Orders []order.Order `json:"orders"`
}

func (ctr Controller) GetOrders(w http.ResponseWriter, r *http.Request) {

	var ordersIdsDto ordersIdsDTO
	err := json.NewDecoder(r.Body).Decode(&ordersIdsDto)
	if err != nil {
		fmt.Println("Error occured while decoding the data: ", err)
		return
	}

	var dto ordersDTO
	dto.Orders, err = ctr.service.GetOrders(ordersIdsDto.OrdersIds)
	if err != nil {
		fmt.Println(err)
		return
	}

	request, err := json.Marshal(dto)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(request)
	return
}
