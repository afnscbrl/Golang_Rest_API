package main

import (
	"fmt"

	"github.com/afnscbrl/Golang_Rest_API/database"
	"github.com/afnscbrl/Golang_Rest_API/routes"
)

func main() {
	//The main only execute the connection with db and get routes
	database.ConnectWithDB()
	fmt.Println("Starting the server...")
	routes.HandleRequest()
}
