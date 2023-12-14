package order

type Order struct {
	OrderId     int64  `json:"order_id"`
	Status      string `json:"status"`
	StoreId     int64  `json:"store_id"`
	DateCreated string `json:"date_created"`
}
