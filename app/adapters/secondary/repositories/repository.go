package repositories

import (
	"database/sql"
	"fmt"
	"github.com/Sachkov98/study/app/adapters/primary/http-adapter/controller"
	"github.com/Sachkov98/study/app/domain/order"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Repository struct {
	dateBase *sql.DB
}

func New() *Repository {
	repository := Repository{}
	return &repository
}

type connectionStrings struct {
	user     string
	password string
	dbname   string
	sslmode  string
}

func (rep *Repository) ConnectToDb() error {

	dbconnectionString := connectionStrings{
		user:     "myUser",
		password: "myPassword",
		dbname:   "postgres",
		sslmode:  "disable"}

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		dbconnectionString.user,
		dbconnectionString.password,
		dbconnectionString.dbname,
		dbconnectionString.sslmode)

	db, err := sql.Open("postgres", connectionString)

	rep.dateBase = db
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (rep Repository) InsertOrders(orders []order.Order) error {

	sqlStatement := `INSERT INTO orders (OrderId, Status, StoreId, DateCreated) VALUES ($1, $2, $3, $4)`

	for _, order := range orders {
		_, err := rep.dateBase.Exec(sqlStatement,
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

type OrdersIds struct {
	OrdersIds []int `json:"orders_ids"`
}

func (rep Repository) GetOrdersByIds(ordersIds controller.OrdersIds) ([]order.Order, error) {

	query := "SELECT * from orders WHERE orderid = ANY ($1)"
	parametrs := pq.Array(ordersIds.OrdersIds)
	rows, err := rep.dateBase.Query(query, parametrs)

	if err != nil {
		return []order.Order{}, err
	}

	var orders []order.Order

	for rows.Next() {
		var orderRow order.Order
		err := rows.Scan(&orderRow.OrderId, &orderRow.Status, &orderRow.StoreId, &orderRow.DateCreated)
		if err != nil {
			return []order.Order{}, err
		}
		orders = append(orders, orderRow)
	}
	defer rows.Close()
	return orders, err
}
