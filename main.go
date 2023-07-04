package main

import (
	"backend-gps-tracker/database"
	"backend-gps-tracker/routes"
)

func main() {
	database.StartDB()

	router := routes.StartServer()

	router.Run()
}
