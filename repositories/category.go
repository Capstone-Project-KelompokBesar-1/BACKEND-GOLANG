package repositories

import (
	"ourgym/models"

	"gorm.io/gorm"
)

func NewCategoryRepository(gormDB *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: gormDB,
	}
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func (cr *CategoryRepositoryImpl) GetAll(name string) []models.Category {
	var categories []models.Category

	cr.db.Find(&categories, "name LIKE ?", "%"+name+"%")

	return categories
}

func (cr *CategoryRepositoryImpl) GetByID(id string) models.Category {
	var category models.Category

	cr.db.First(&category, id)

	return category
}
