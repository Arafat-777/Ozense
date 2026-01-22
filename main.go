package main

import (
	"ozenshe/config"
	"ozenshe/routes"
)

func main() {
	config.InitDB()
	r := routes.SetupRoutes()
	r.Run(":8080")
}
