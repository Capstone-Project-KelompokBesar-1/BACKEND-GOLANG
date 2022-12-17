package routes

import (
	"net/http"
	"ourgym/config"
	"ourgym/controllers"
	"ourgym/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	AuthController          controllers.AuthController
	UserController          controllers.UserController
	ProfileController       controllers.ProfileController
	ClassController         controllers.ClassController
	TrainerController       controllers.TrainerController
	TransactionController   controllers.TransactionController
	PaymentMethodController controllers.PaymentMethodController
	MeController            controllers.MeController
	CategoryController      controllers.CategoryController
	DashboardController     controllers.DashboardController
}

func (cl ControllerList) InitRoute() *echo.Echo {
	e := echo.New()
	cfg := config.Cfg

	configAdmin := middleware.JWTConfig{
		KeyFunc: middlewares.GetJWTSecretKeyForAdmin,
	}

	adminJwtMiddleware := middleware.JWTWithConfig(configAdmin)
	userJwtMiddleware := middleware.JWT([]byte(cfg.JWT_SECRET_KEY))

	e.POST("/login", cl.AuthController.Login)
	e.POST("/register", cl.AuthController.Register)
	e.POST("/forgot-password", cl.AuthController.ForgotPassword)
	e.POST("/validate-otp", cl.AuthController.ValidateOTP)
	e.POST("/reset-password", cl.AuthController.ResetPassword, userJwtMiddleware)

	user := e.Group("/user")

	user.GET("/profile", cl.ProfileController.GetProfile, userJwtMiddleware)
	user.PUT("/profile", cl.ProfileController.UpdateProfile, userJwtMiddleware)
	user.PUT("/change-password", cl.ProfileController.ChangePassword, userJwtMiddleware)
	user.POST("user/refresh-token", cl.AuthController.RefreshToken, userJwtMiddleware)
	user.GET("/my-class-online/:id", cl.MeController.OnlineClass, userJwtMiddleware)
	user.GET("/my-class-offline/:id", cl.MeController.OfflineClass, userJwtMiddleware)

	users := e.Group("/users")

	users.GET("", cl.UserController.GetAll, adminJwtMiddleware)
	users.GET("/:id", cl.UserController.GetByID, adminJwtMiddleware)
	users.POST("", cl.UserController.Create, adminJwtMiddleware)
	users.PUT("/:id", cl.UserController.Update, adminJwtMiddleware)
	users.DELETE("/:id", cl.UserController.Delete, adminJwtMiddleware)
	users.DELETE("", cl.UserController.DeleteMany, adminJwtMiddleware)

	classes := e.Group("/classes")

	classes.GET("", cl.ClassController.GetAll, userJwtMiddleware)
	classes.GET("/online", cl.ClassController.GetAllOnlineClass, userJwtMiddleware)
	classes.GET("/offline", cl.ClassController.GetAllOfflineClass, userJwtMiddleware)
	classes.GET("/:id", cl.ClassController.GetByID, userJwtMiddleware)
	classes.POST("", cl.ClassController.Create, adminJwtMiddleware)
	classes.PUT("/:id", cl.ClassController.Update, adminJwtMiddleware)
	classes.DELETE("/:id", cl.ClassController.Delete, adminJwtMiddleware)
	classes.DELETE("", cl.ClassController.DeleteMany, adminJwtMiddleware)

	trainers := e.Group("/trainers")

	trainers.GET("", cl.TrainerController.GetAll, userJwtMiddleware)
	trainers.GET("/:id", cl.TrainerController.GetByID, userJwtMiddleware)

	paymentMethods := e.Group("/payment-methods")

	paymentMethods.GET("", cl.PaymentMethodController.GetAll, userJwtMiddleware)
	paymentMethods.GET("/:id", cl.PaymentMethodController.GetByID, userJwtMiddleware)

	transactions := e.Group("/transactions")
	transactions.GET("", cl.TransactionController.GetAll, adminJwtMiddleware)
	transactions.GET("/history", cl.TransactionController.GetHistory, adminJwtMiddleware)
	transactions.GET("/user/:id", cl.TransactionController.GetByUserID, userJwtMiddleware)
	transactions.GET("/:id", cl.TransactionController.GetByID, userJwtMiddleware)
	transactions.POST("", cl.TransactionController.Create, userJwtMiddleware)
	transactions.POST("/midtrans-api", cl.TransactionController.UpdatedByMidtransAPI)
	transactions.PUT("/:id", cl.TransactionController.Update, adminJwtMiddleware)
	transactions.DELETE("/:id", cl.TransactionController.Delete, adminJwtMiddleware)
	transactions.DELETE("", cl.TransactionController.DeleteMany, adminJwtMiddleware)

	categories := e.Group("/categories")

	categories.GET("", cl.CategoryController.GetAll, userJwtMiddleware)
	categories.GET("/:id", cl.CategoryController.GetByID, userJwtMiddleware)

	dashboard := e.Group("/dashboard")
	dashboard.GET("", cl.DashboardController.GetData, adminJwtMiddleware)

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get data",
			"data":    "Hello world!",
		})
	})

	return e
}
