package main

import (
	"os"
	"ourgym/config"
	"ourgym/controllers"
	"ourgym/controllers/admin/manage/users"
	"ourgym/controllers/user"
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
	userService := services.NewUserService(userRepo)
	userController := users.NewUserController(userService)

	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	profileController := user.NewProfileController(userService)

	route := routes.ControllerList{
		AuthController:    *authController,
		UserController:    *userController,
		ProfileController: *profileController,
	}

	e := route.InitRoute()

	middlewares.Logger(e)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
