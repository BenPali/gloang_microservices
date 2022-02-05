package main

import (
	"fmt"

	service "github.com/k4lii/golang_microservices/ads/cmd/service"
	"github.com/k4lii/golang_microservices/db/connection"
)

func main() {
	err := connection.InitDB()
	if err != nil {
		fmt.Println("Could not connect to the database:", err)
	} else {
		service.Run()
	}
}
