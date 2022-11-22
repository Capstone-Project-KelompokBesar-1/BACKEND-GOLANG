package main

import (
	"ourgym/config"
	"ourgym/routes"
)

func main() {
	config.InitConfig()

	route := routes.InitRoute()

	route.Logger.Fatal(route.Start(":8080"))
}
