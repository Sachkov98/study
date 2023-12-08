package repositories

import (
	"database/sql"
	"fmt"
	_ "github.codatabase/sqlm/lib/pq"
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

func (rep repository) ConnectToDataBase() (sql.DB, error){

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

func (rep repository) InsertListOrdersToDb() {

	sqlStatement := `INSERT INTO list_orders (OrderId, Status, StoreId, DateCreated) VALUES ($1, $2, $3, $4)`

	order := gateways.New().GetBody

	_, err = .Exec(sqlStatement,
		order.OrderId,
		order.Status,
		order.StoreId,
		order.DateCreated)
	if err != nil {
		log.Fatal(err)
	}
}
