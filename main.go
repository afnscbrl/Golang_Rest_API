package main

import (
	"fmt"

	"github.com/afnscbrl/Golang_Rest_API/database"
	"github.com/afnscbrl/Golang_Rest_API/routes"
)

func main() {
	database.ConnectWithDB()
	fmt.Println("Starting the server...")
	routes.HandleRequest()
}
