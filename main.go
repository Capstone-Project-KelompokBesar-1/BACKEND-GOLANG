package main

import (
	"os"
	"ourgym/config"
	"ourgym/controllers"
	"ourgym/controllers/admin/manage/users"
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
	userService := services.NewUserService(userRepo)
	authController := controllers.NewAuthController(authService)
	adminUserController := users.NewAdminUserController(userService)

	route := routes.ControllerList{
		AuthController:      *authController,
		AdminUserController: *adminUserController,
	}

	e := route.InitRoute()

	middlewares.Logger(e)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
