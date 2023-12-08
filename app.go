package main

import (
	"study/app/adapters/secondary/gateways"
	"study/app/adapters/secondary/repositories"
	"study/app/services"
)

func main() {
	services.New(repositories.New(), gateways.New()).Start()
}
