package repositories

import (
	"ourgym/models"

	"gorm.io/gorm"
)

func NewOtpRepository(gormDB *gorm.DB) OtpRepository {
	return &OtpRepositoryImpl{
		db: gormDB,
	}
}

type OtpRepositoryImpl struct {
	db *gorm.DB
}

func (otpR *OtpRepositoryImpl) GetOneByFilter(key string, value any) models.Otp {
	var otp models.Otp
	otpR.db.First(&otp, key, value)

	return otp
}
