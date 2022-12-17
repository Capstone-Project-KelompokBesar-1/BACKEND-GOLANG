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

func (otpR *OtpRepositoryImpl) Create(otp models.Otp) models.Otp {
	rec := otpR.db.Create(&otp)

	rec.Last(&otp)

	return otp
}

func (otpR *OtpRepositoryImpl) Delete(id uint) bool {
	otp := otpR.GetOneByFilter("id", id)

	rec := otpR.db.Delete(&otp)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
