package main

import (
	"gps-tracker/database"
	"gps-tracker/routes"
)

func main() {
	database.StartDB()

	router := routes.StartServer()

	router.Run()
}
