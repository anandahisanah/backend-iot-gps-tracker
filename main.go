package main

import (
	"backend-iot-gps-tracker/database"
	"backend-iot-gps-tracker/routes"
)

func main() {
	database.StartDB()

	router := routes.StartServer()

	router.Run()
}
