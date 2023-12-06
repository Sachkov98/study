package main

import (
	"study/app/services"
)

func main() {
	services.New(services.Intf).GetNDrop()
}
