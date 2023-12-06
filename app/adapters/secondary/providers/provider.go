package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type provider struct {
}

func New() *provider {
	prov := provider{}
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

var Dto DTO

func (p provider) GetBody() {

	resp, err := http.Get("http://localhost:8081")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &Dto)
	if err != nil {
		log.Fatal(err)
	}

}
