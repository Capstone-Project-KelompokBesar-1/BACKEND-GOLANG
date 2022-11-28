package main

import (
	"os"
	"ourgym/config"
	"ourgym/controllers"
	"ourgym/databases"
	"ourgym/middlewares"
	"ourgym/repositories"
	"ourgym/routes"
	"ourgym/services"
)

func main() {
	config.InitConfig()
	db := databases.InitDatabase()

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	route := routes.ControllerList{
		AuthController: *authController,
	}

	e := route.InitRoute()

	middlewares.Logger(e)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
