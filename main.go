package main

import (
	"ourgym/routes"
)

func main() {
	route := routes.InitRoute()
	route.Logger.Fatal(route.Start(":8080"))
}
