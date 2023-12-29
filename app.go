package main

import (
	"github.com/Sachkov98/study/app/adapters/primary/http-adapter/controller"
	"github.com/Sachkov98/study/app/adapters/secondary/gateways"
	"github.com/Sachkov98/study/app/adapters/secondary/repositories"
	"github.com/Sachkov98/study/app/services"
	"log"
	"net/http"
)

func main() {

	gateway := gateways.New()
	repository := repositories.New()
	service := services.New(gateway, repository)
	controller := controller.New(service)
	err := repository.ConnectToDb()
	if err != nil {
		log.Fatal()
	}

	go service.Start()

	http.HandleFunc("/orders", controller.GetOrders)

	log.Println("Listening...")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal()
	}

}
