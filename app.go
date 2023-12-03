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

type orders struct {
	Content []struct {
		OrderId     int64  `json:"order_id"`
		Status      string `json:"status"`
		StoreId     int64  `json:"store_id"`
		DateCreated string `json:"date_created"`
	} `json:"content"`
}

type connectionString struct {
	user, password, dbname, sslmode string
}

func main() {

	for {

		var (
			orders1           orders
			connectionString1 = connectionString{
				user:     "myUser",
				password: "myPassword",
				dbname:   "postgres",
				sslmode:  "disable"}
		)

		resp, err := http.Get("http://localhost:8081")

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(body, &orders1)

		if err != nil {

			log.Fatal(err)
		}

		//connectionString := "user=myUser password=myPassword dbname=postgres sslmode=disable"
		db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			connectionString1.user,
			connectionString1.password,
			connectionString1.dbname,
			connectionString1.sslmode))

		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Успешное подключение к БД")
		defer db.Close()

		sqlStatement := `INSERT INTO list (OrderId, Status, StoreId, DateCreated) VALUES ($1, $2, $3, $4)`
		_, err = db.Exec(sqlStatement,
			orders1.Content[0].OrderId,
			orders1.Content[0].Status,
			orders1.Content[0].StoreId,
			orders1.Content[0].DateCreated)

		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nRow inserted successfully!")
		}
		time.Sleep(60 * time.Second)

	}
}
