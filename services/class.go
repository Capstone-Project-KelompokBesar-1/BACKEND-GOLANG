package services

import (
	"ourgym/dto"
	"ourgym/models"
	"ourgym/repositories"
)

func NewClassService(cr repositories.ClassRepository) ClassService {
	return &ClassServiceImpl{
		classRepository: cr,
	}
}

type ClassServiceImpl struct {
	classRepository repositories.ClassRepository
}

func (cs *ClassServiceImpl) GetAll(classType string, name string) []dto.ClassResponse {
	classes := cs.classRepository.GetAll(classType, name)

	var classesResponse []dto.ClassResponse

	for _, class := range classes {
		classesResponse = append(classesResponse, class.ConvertToDTO())
	}

	return classesResponse
}

func (cs *ClassServiceImpl) GetByID(id string) dto.ClassResponse {
	class := cs.classRepository.GetOneByFilter("id", id)
	return class.ConvertToDTO()
}

func (cs *ClassServiceImpl) Create(classRequest dto.ClassRequest) dto.ClassResponse {
	classModel := models.FromClassRequestToClassModel(classRequest)
	class := cs.classRepository.Create(classModel)
	return class.ConvertToDTO()
}

func (cs *ClassServiceImpl) Update(id string, classRequest dto.ClassRequest) dto.ClassResponse {
	classModel := models.FromClassRequestToClassModel(classRequest)
	class := cs.classRepository.Update(id, classModel)
	return class.ConvertToDTO()
}

func (cs *ClassServiceImpl) Delete(id string) bool {
	return cs.classRepository.Delete(id)
}

func (cs *ClassServiceImpl) DeleteMany(ids string) bool {
	return cs.classRepository.DeleteMany(ids)
}
