// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "ourgym/models"

	mock "github.com/stretchr/testify/mock"
)

// ArticleRepository is an autogenerated mock type for the ArticleRepository type
type ArticleRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: articleRequest
func (_m *ArticleRepository) Create(articleRequest models.Article) models.Article {
	ret := _m.Called(articleRequest)

	var r0 models.Article
	if rf, ok := ret.Get(0).(func(models.Article) models.Article); ok {
		r0 = rf(articleRequest)
	} else {
		r0 = ret.Get(0).(models.Article)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *ArticleRepository) Delete(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeleteManyArticle provides a mock function with given fields: ids
func (_m *ArticleRepository) DeleteManyArticle(ids string) bool {
	ret := _m.Called(ids)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(ids)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields: title
func (_m *ArticleRepository) GetAll(title string) []models.Article {
	ret := _m.Called(title)

	var r0 []models.Article
	if rf, ok := ret.Get(0).(func(string) []models.Article); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Article)
		}
	}

	return r0
}

// GetArticleByID provides a mock function with given fields: articleID
func (_m *ArticleRepository) GetArticleByID(articleID string) []models.Article {
	ret := _m.Called(articleID)

	var r0 []models.Article
	if rf, ok := ret.Get(0).(func(string) []models.Article); ok {
		r0 = rf(articleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Article)
		}
	}

	return r0
}

// Update provides a mock function with given fields: id, articleRequest
func (_m *ArticleRepository) Update(id string, articleRequest models.Article) models.Article {
	ret := _m.Called(id, articleRequest)

	var r0 models.Article
	if rf, ok := ret.Get(0).(func(string, models.Article) models.Article); ok {
		r0 = rf(id, articleRequest)
	} else {
		r0 = ret.Get(0).(models.Article)
	}

	return r0
}

type mockConstructorTestingTNewArticleRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewArticleRepository creates a new instance of ArticleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewArticleRepository(t mockConstructorTestingTNewArticleRepository) *ArticleRepository {
	mock := &ArticleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
