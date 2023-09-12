package main

import (
	"it15th/database"
	"it15th/routes"
)

func main() {
	database.ConnectDB()
	routes.SetupRoute()
}
