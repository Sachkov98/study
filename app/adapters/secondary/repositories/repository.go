package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"study/app/domain/order"
)

type Repository struct {
	dateBase *sql.DB
}

func New() *Repository {
	repo := Repository{}
	return &repo
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
