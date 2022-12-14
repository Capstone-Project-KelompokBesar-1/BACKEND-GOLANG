package repositories

import (
	"ourgym/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewMeRepository(gormDB *gorm.DB) MeRepository {
	return &MeRepositoryImpl{
		db: gormDB,
	}
}

type MeRepositoryImpl struct {
	db *gorm.DB
}

func (me *MeRepositoryImpl) GetTransactionByID(userId string) []models.Transaction {
	var transactions []models.Transaction

	me.db.Preload(clause.Associations).Find(&transactions, "user_id = ?", userId)

	return transactions
}

// func (me *MeRepositoryImpl) OnlineClass(userId string) []models.Transaction {

// 	return []models.Transaction{}
// }

// func (me *MeRepositoryImpl) OfflineClass(userId string) []models.Transaction {

// 	return []models.Transaction{}
// }
