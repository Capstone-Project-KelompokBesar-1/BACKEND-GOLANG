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
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	classRepo := repositories.NewClassRepository(db)
	classService := services.NewClassService(classRepo)
	classController := controllers.NewClassController(classService)

	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	profileController := controllers.NewProfileController(userService)

	route := routes.ControllerList{
		AuthController:    *authController,
		UserController:    *userController,
		ProfileController: *profileController,
		ClassController:   *classController,
	}

	e := route.InitRoute()

	middlewares.Logger(e)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
