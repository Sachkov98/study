package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"study/app/adapters/secondary/gateways"
)

type repository struct{}

func New() *repository {
	repo := repository{}
	return &repo
}

type connectionStrings struct {
	user     string
	password string
	dbname   string
	sslmode  string
}

func (rep repository) ConnectToDb() (*sql.DB, error) {

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
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()
	return db, err
}

func (rep repository) InsertOrdersToDb(ord []gateways.Order, db *sql.DB) error {

	sqlStatement := `INSERT INTO orders (OrderId, Status, StoreId, DateCreated) VALUES ($1, $2, $3, $4)`

	orders := ord[0]
	_, err := db.Exec(sqlStatement,
		orders.OrderId,
		orders.Status,
		orders.StoreId,
		orders.DateCreated)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
