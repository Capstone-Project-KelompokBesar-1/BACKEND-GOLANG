package services

import (
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

func (cs *ClassServiceImpl) GetAll(classType string, name string) []models.Class {
	return cs.classRepository.GetAll(classType, name)
}

func (cs *ClassServiceImpl) GetByID(id string) models.Class {
	return cs.classRepository.GetOneByFilter("id", id)
}

func (cs *ClassServiceImpl) Create(classRequest models.Class) models.Class {
	return cs.classRepository.Create(classRequest)
}

func (cs *ClassServiceImpl) Update(id string, classRequest models.Class) models.Class {
	return cs.classRepository.Update(id, classRequest)
}

func (cs *ClassServiceImpl) UpdatePhoto(id string, classRequest models.Class) models.Class {
	return cs.classRepository.Update(id, classRequest)
}

func (cs *ClassServiceImpl) Delete(id string) bool {
	return cs.classRepository.Delete(id)
}

func (cs *ClassServiceImpl) DeleteMany(ids string) bool {
	return cs.classRepository.DeleteMany(ids)
}
