package repositories

import (
	"ourgym/models"

	"gorm.io/gorm"
)

func NewTrainerRepository(gormDB *gorm.DB) TrainerRepository {
	return &TrainerRepositoryImpl{
		db: gormDB,
	}
}

type TrainerRepositoryImpl struct {
	db *gorm.DB
}

func (tr *TrainerRepositoryImpl) GetAll(name string) []models.Trainer {
	var trainers []models.Trainer

	tr.db.Find(&trainers, "name LIKE ?", "%"+name+"%")

	return trainers
}

func (tr *TrainerRepositoryImpl) GetByID(id string) models.Trainer {
	var trainer models.Trainer

	tr.db.First(&trainer, id)

	return trainer
}

func (tr *TrainerRepositoryImpl) CountTrainer() int64 {
	var total int64

	tr.db.Find(&models.Trainer{}).Count(&total)

	return total
}
