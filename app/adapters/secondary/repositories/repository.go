package repositories

import (
	"database/sql"
	"fmt"
	"github.com/Sachkov98/study/app/domain/order"
	"github.com/lib/pq"
)

type Repository struct {
	dataBase *sql.DB
}

func New() *Repository {
	repository := Repository{}
	return &repository
}

type config struct {
	user     string
	password string
	dbname   string
	sslmode  string
}

func (rep *Repository) ConnectToDb() error {
	config := config{
		user:     "myUser",
		password: "myPassword",
		dbname:   "postgres",
		sslmode:  "disable"}

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		config.user,
		config.password,
		config.dbname,
		config.sslmode)

	dataBase, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	err = dataBase.Ping()
	if err != nil {
		return err
	}

	rep.dataBase = dataBase
	return nil
}

func (rep Repository) InsertOrders(orders []order.Order) error {
	sqlStatement := `INSERT INTO orders (OrderId, Status, StoreId, DateCreated) VALUES ($1, $2, $3, $4)`
	for _, order := range orders {
		_, err := rep.dataBase.Exec(sqlStatement,
			order.OrderId,
			order.Status,
			order.StoreId,
			order.DateCreated)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rep Repository) GetOrdersByIds(ordersIds []int) ([]order.Order, error) {
	query := "SELECT * from orders WHERE orderid = ANY ($1)"
	parametrs := pq.Array(ordersIds)
	rows, err := rep.dataBase.Query(query, parametrs)
	if err != nil {
		return []order.Order{}, err
	}

	var orders []order.Order
	for rows.Next() {
		var order order.Order
		err := rows.Scan(&order.OrderId, &order.Status, &order.StoreId, &order.DateCreated)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	defer rows.Close()
	return orders, nil
}
