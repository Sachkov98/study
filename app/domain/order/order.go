package order

type Order struct {
	OrderID     int64  `json:"order_id"`
	Status      string `json:"status"`
	StoreID     int64  `json:"store_id"`
	DateCreated string `json:"date_created"`
}
