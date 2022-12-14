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

	classRepo := repositories.NewClassRepository(db)
	otpRepo := repositories.NewOtpRepository(db)
	paymentMethodRepo := repositories.NewPaymentMethodRepository(db)
	trainerRepo := repositories.NewTrainerRepository(db)
	userRepo := repositories.NewUserRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db, userRepo, paymentMethodRepo)
	MeRepo := repositories.NewMeRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	authService := services.NewAuthService(userRepo, otpRepo)
	classService := services.NewClassService(classRepo)
	paymentMethodService := services.NewPaymentMethodService(paymentMethodRepo)
	trainerService := services.NewTrainerService(trainerRepo)
	transactionService := services.NewTransactionService(transactionRepo)
	userService := services.NewUserService(userRepo)
	meService := services.NewMeService(MeRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	dashboardService := services.NewDashboardService(userRepo, trainerRepo, classRepo, transactionRepo)

	authController := controllers.NewAuthController(authService)
	classController := controllers.NewClassController(classService)
	profileController := controllers.NewProfileController(userService)
	paymentMethodController := controllers.NewPaymentMethodController(paymentMethodService)
	trainerController := controllers.NewTrainerController(trainerService)
	transactionController := controllers.NewTransactionController(transactionService)
	userController := controllers.NewUserController(userService)
	meController := controllers.NewMeController(meService)
	categoryController := controllers.NewCategoryController(categoryService)
	dashboardController := controllers.NewDashboardController(dashboardService)

	route := routes.ControllerList{
		AuthController:          *authController,
		UserController:          *userController,
		ProfileController:       *profileController,
		ClassController:         *classController,
		TrainerController:       *trainerController,
		TransactionController:   *transactionController,
		PaymentMethodController: *paymentMethodController,
		MeController:            *meController,
		CategoryController:      *categoryController,
		DashboardController:     *dashboardController,
	}

	e := route.InitRoute()

	middlewares.CORS(e)

	middlewares.Logger(e)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
