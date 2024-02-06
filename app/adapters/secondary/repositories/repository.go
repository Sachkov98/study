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
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func (rep *Repository) ConnectToDB() error {
	config := config{
		Host:     "postgres_db",
		Port:     5432,
		User:     "myUser",
		Password: "myPassword",
		DbName:   "postgres",
	}

	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DbName)

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
	sqlStatement := `INSERT INTO orders (OrderID, Status, StoreID, DateCreated) VALUES ($1, $2, $3, $4)`
	for _, order := range orders {
		_, err := rep.dataBase.Exec(sqlStatement,
			order.OrderID,
			order.Status,
			order.StoreID,
			order.DateCreated)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rep Repository) GetOrdersByIds(ordersIds []int) ([]order.Order, error) {
	parametrs := pq.Array(ordersIds)
	query := "SELECT * from orders WHERE orderID = ANY ($1)"

	rows, err := rep.dataBase.Query(query, parametrs)
	if err != nil {
		return []order.Order{}, err
	}

	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			if err == nil {
				err = closeErr
			}
		}
	}()

	var orders []order.Order

	for rows.Next() {
		var order order.Order

		err := rows.Scan(&order.OrderID, &order.Status, &order.StoreID, &order.DateCreated)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
