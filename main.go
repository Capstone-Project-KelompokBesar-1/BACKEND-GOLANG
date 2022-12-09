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
  
	otpRepo := repositories.NewOtpRepository(db)
	authService := services.NewAuthService(userRepo, otpRepo)
  
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	classRepo := repositories.NewClassRepository(db)
	classService := services.NewClassService(classRepo)
	classController := controllers.NewClassController(classService)

	authController := controllers.NewAuthController(authService)

	profileController := controllers.NewProfileController(userService)

	trainerRepo := repositories.NewTrainerRepository(db)
	trainerService := services.NewTrainerService(trainerRepo)
	trainerController := controllers.NewTrainerController(trainerService)

	route := routes.ControllerList{
		AuthController:    *authController,
		UserController:    *userController,
		ProfileController: *profileController,
		ClassController:   *classController,
		TrainerController: *trainerController,
	}

	e := route.InitRoute()

	middlewares.Logger(e)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
