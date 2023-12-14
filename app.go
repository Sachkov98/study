package main

import (
	"log"
	"study/app/adapters/secondary/gateways"
	"study/app/adapters/secondary/repositories"
	"study/app/services"
)

func main() {
	gateway := gateways.New()
	repository := repositories.New()
	service := services.New(gateway, repository)

	err := repository.ConnectToDb()
	if err != nil {
		log.Fatal()
	}

	service.Start()

}
