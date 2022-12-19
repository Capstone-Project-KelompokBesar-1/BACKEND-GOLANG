package repositories

import (
	"ourgym/models"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewClassRepository(gormDB *gorm.DB) ClassRepository {
	return &ClassRepositoryImpl{
		db: gormDB,
	}
}

type ClassRepositoryImpl struct {
	db *gorm.DB
}

func (cr *ClassRepositoryImpl) GetAll(classType string, name string) []models.Class {
	var class []models.Class

	if classType == "" {
		cr.db.Preload(clause.Associations).Find(&class, "name LIKE ?", "%"+name+"%")
	} else {
		cr.db.Preload(clause.Associations).Find(&class, "type = ? && name LIKE ?", classType, "%"+name+"%")
	}

	return class
}

func (cr *ClassRepositoryImpl) GetOneByFilter(key string, value any) models.Class {
	var class models.Class

	cr.db.Preload(clause.Associations).First(&class, key, value)

	return class
}

func (cr *ClassRepositoryImpl) Create(classRequest models.Class) models.Class {
	var class models.Class

	rec := cr.db.Create(&classRequest)

	rec.Preload(clause.Associations).Last(&class)

	return class
}

func (cr *ClassRepositoryImpl) Update(id string, classRequest models.Class) models.Class {
	class := cr.GetOneByFilter("id", id)

	class.Name = classRequest.Name
	class.Thumbnail = classRequest.Thumbnail
	class.TrainerID = classRequest.TrainerID
	class.Trainer.ID = classRequest.TrainerID
	class.Description = classRequest.Description
	class.Type = classRequest.Type
	class.Category.ID = classRequest.CategoryID
	class.Price = classRequest.Price
	class.TotalMeeting = classRequest.TotalMeeting

	rec := cr.db.Save(&class)

	rec.Preload(clause.Associations).Last(&class)

	return class
}

func (cr *ClassRepositoryImpl) Delete(id string) bool {
	class := cr.GetOneByFilter("id", id)

	rec := cr.db.Select("Transactions").Delete(&class)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}

func (cr *ClassRepositoryImpl) DeleteMany(ids string) bool {
	classIds := strings.Split(ids, ",")

	var classes []models.Class

	cr.db.Find(&classes, "id IN (?)", classIds)

	rec := cr.db.Select("Transactions").Delete(&classes, "id IN (?)", classIds)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}

func (cr *ClassRepositoryImpl) CountClass() int64 {
	var total int64

	cr.db.Find(&models.Class{}).Count(&total)

	return total
}
