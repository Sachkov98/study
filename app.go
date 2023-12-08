package main

import (
	"fmt"
	"study/app/adapters/secondary/gateways"
)

func main() {
	fmt.Println(gateways.New().GetOrders())
}
