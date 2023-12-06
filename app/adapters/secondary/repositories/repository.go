package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"study/app/adapters/secondary/providers"
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

func (rep repository) InsertTable() {

	DBconnectionString := connectionStrings{
		user:     "myUser",
		password: "myPassword",
		dbname:   "postgres",
		sslmode:  "disable"}

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		DBconnectionString.user,
		DBconnectionString.password,
		DBconnectionString.dbname,
		DBconnectionString.sslmode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStatement := `INSERT INTO list_orders (OrderId, Status, StoreId, DateCreated) VALUES ($1, $2, $3, $4)`

	order := providers.Dto.Orders[0]

	_, err = db.Exec(sqlStatement,
		order.OrderId,
		order.Status,
		order.StoreId,
		order.DateCreated)
	if err != nil {
		log.Fatal(err)
	}
}
