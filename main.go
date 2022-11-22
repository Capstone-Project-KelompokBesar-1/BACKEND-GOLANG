package main

import (
	"os"
	"ourgym/config"
	"ourgym/routes"
)

func main() {
	config.InitConfig()

	route := routes.InitRoute()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	route.Logger.Fatal(route.Start(":" + port))
}
