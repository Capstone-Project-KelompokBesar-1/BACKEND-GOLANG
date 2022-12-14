package services

import (
	"ourgym/dto"
	"ourgym/repositories"
)

func NewCategoryService(cr repositories.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: cr,
	}
}

type CategoryServiceImpl struct {
	categoryRepository repositories.CategoryRepository
}

func (cs *CategoryServiceImpl) GetAll(name string) []dto.CategoryResponse {
	categories := cs.categoryRepository.GetAll(name)

	var categoriesResponse []dto.CategoryResponse

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, category.ConvertToDTO())
	}

	return categoriesResponse
}

func (cs *CategoryServiceImpl) GetByID(id string) dto.CategoryResponse {
	category := cs.categoryRepository.GetByID(id)
	return category.ConvertToDTO()
}
