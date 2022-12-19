package databases

import (
	"fmt"
	"ourgym/config"
	"ourgym/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	cfg := config.Cfg

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Otp{},
		&models.Class{},
		&models.Trainer{},
		&models.PaymentMethod{},
		&models.Transaction{},
	)

	return db

}
