package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Order struct {
	OrderId     int64  `json:"order_id"`
	Status      string `json:"status"`
	StoreId     int64  `json:"store_id"`
	DateCreated string `json:"date_created"`
}

type DTO struct {
	Orders []Order `json:"content"`
}

type connectionStrings struct {
	user     string
	password string
	dbname   string
	sslmode  string
}

func main() {

	for {
		var dto DTO

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

		resp, err := http.Get("http://localhost:8081")
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(body, &dto)
		if err != nil {
			log.Fatal(err)
		}

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

		_, err = db.Exec(sqlStatement,
			dto.Orders[0].OrderId,
			dto.Orders[0].Status,
			dto.Orders[0].StoreId,
			dto.Orders[0].DateCreated)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(60 * time.Second)
	}
}
