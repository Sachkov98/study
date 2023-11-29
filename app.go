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

func main() {
	rest_indus()
}

// add
func rest_indus() {

	type T struct {
		Content []struct {
			OrderId     int64  `json:"order_id"`
			Status      string `json:"status"`
			StoreId     int64  `json:"store_id"`
			DateCreated string `json:"date_created"`
		} `json:"content"`
	}

	var t1 T

	resp, err := http.Get("http://localhost:8081")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &t1)

	if err != nil {

		log.Fatal(err)
	}
	fmt.Println(t1)
	fmt.Println(t1.Content)

	connStr := "user=myUser password=myPassword dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
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
	_, err = db.Exec(sqlStatement, t1.Content[0].OrderId, t1.Content[0].Status, t1.Content[0].StoreId, t1.Content[0].DateCreated)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
	time.Sleep(60 * time.Second)

}
