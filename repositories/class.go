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

func (ur *ClassRepositoryImpl) GetAll(classType string, name string) []models.Class {
	var class []models.Class

	if classType == "" {
		ur.db.Preload(clause.Associations).Find(&class, "name LIKE ?", "%"+name+"%")
	} else {
		ur.db.Preload(clause.Associations).Find(&class, "type = ? && name LIKE ?", classType, "%"+name+"%")
	}

	return class
}

func (ur *ClassRepositoryImpl) GetOneByFilter(key string, value any) models.Class {
	var class models.Class

	ur.db.Preload(clause.Associations).First(&class, key, value)

	return class
}

func (ur *ClassRepositoryImpl) Create(classRequest models.Class) models.Class {
	var class models.Class

	rec := ur.db.Create(&classRequest)

	rec.Preload(clause.Associations).Last(&class)

	return class
}

func (ur *ClassRepositoryImpl) Update(id string, classRequest models.Class) models.Class {
	class := ur.GetOneByFilter("id", id)

	class.Name = classRequest.Name
	class.Thumbnail = classRequest.Thumbnail
	class.TrainerID = classRequest.TrainerID
	class.Trainer.ID = classRequest.TrainerID
	class.Description = classRequest.Description
	class.Type = classRequest.Type
	class.Category.ID = classRequest.CategoryID

	rec := ur.db.Save(&class)

	rec.Preload(clause.Associations).Last(&class)

	return class
}

func (ur *ClassRepositoryImpl) Delete(id string) bool {
	class := ur.GetOneByFilter("id", id)

	rec := ur.db.Delete(&class)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}

func (ur *ClassRepositoryImpl) DeleteMany(ids string) bool {
	classIds := strings.Split(ids, ",")

	rec := ur.db.Delete(&models.Class{}, "id IN (?)", classIds)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
