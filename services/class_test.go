package services

import (
	"ourgym/dto"
	"ourgym/models"
	"ourgym/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type suiteClass struct {
	suite.Suite
	service       ClassService
	classRepoMock *mocks.ClassRepository
	otpRepoMock   *mocks.OtpRepository
	classRequest  dto.ClassRequest
	classModel    models.Class
}

func (s *suiteClass) SetupSuite() {
	s.classRepoMock = &mocks.ClassRepository{}

	s.service = NewClassService(s.classRepoMock)

	s.classRequest = dto.ClassRequest{
		TrainerID:    1,
		CategoryID:   1,
		Name:         "class testing",
		Description:  "test description",
		TotalMeeting: 4,
		Thumbnail:    "http://",
		Type:         "offline",
		Price:        99000,
	}

	s.classModel = models.Class{
		ID:           1,
		TrainerID:    1,
		CategoryID:   1,
		Name:         "class testing",
		Description:  "test description",
		TotalMeeting: 4,
		Thumbnail:    "http://",
		Type:         "offline",
		Price:        99000,
	}

}

func (s *suiteClass) TestGetAll() {
	s.T().Run("GetAll | Valid", func(t *testing.T) {
		s.classRepoMock.On("GetAll", "", "").Return([]models.Class{s.classModel}).Once()

		classes := s.service.GetAll("", "")

		assert.Equal(t, 1, len(classes))
	})

	s.T().Run("GetAll | Invalid", func(t *testing.T) {
		s.classRepoMock.On("GetAll", "", "").Return([]models.Class{}).Once()

		classes := s.service.GetAll("", "")

		assert.Equal(t, 0, len(classes))
	})
}

func (s *suiteClass) TestGetByID() {
	s.T().Run("GetByID | Valid", func(t *testing.T) {
		s.classRepoMock.On("GetOneByFilter", "id", "1").Return(s.classModel).Once()

		classes := s.service.GetByID("1")

		assert.Equal(t, uint(1), classes.ID)
	})

	s.T().Run("GetByID | Invalid", func(t *testing.T) {
		s.classRepoMock.On("GetOneByFilter", "id", "").Return(models.Class{}).Once()

		classes := s.service.GetByID("")

		assert.Equal(t, uint(0), classes.ID)
	})
}

func (s *suiteClass) TestCreate() {
	s.T().Run("Create | Valid", func(t *testing.T) {
		classReq := models.FromClassRequestToClassModel(s.classRequest)

		s.classRepoMock.On("Create", classReq).Return(s.classModel).Once()

		classes := s.service.Create(s.classRequest)

		assert.Equal(t, uint(1), classes.ID)
	})

	s.T().Run("Create | Invalid", func(t *testing.T) {
		s.classRepoMock.On("Create", models.Class{}).Return(models.Class{}).Once()

		classes := s.service.Create(dto.ClassRequest{})

		assert.Equal(t, uint(0), classes.ID)
	})
}

func (s *suiteClass) TestUpdate() {
	s.T().Run("Update | Valid", func(t *testing.T) {
		classReq := models.FromClassRequestToClassModel(s.classRequest)

		s.classRepoMock.On("Update", "1", classReq).Return(s.classModel).Once()

		classes := s.service.Update("1", s.classRequest)

		assert.Equal(t, uint(1), classes.ID)
	})

	s.T().Run("Update | Invalid", func(t *testing.T) {
		s.classRepoMock.On("Update", "0", models.Class{}).Return(models.Class{}).Once()

		classes := s.service.Update("0", dto.ClassRequest{})

		assert.Equal(t, uint(0), classes.ID)
	})
}

func (s *suiteClass) TestDelete() {
	s.T().Run("Delete | Valid", func(t *testing.T) {

		s.classRepoMock.On("Delete", "1").Return(true).Once()

		isSuccess := s.service.Delete("1")

		assert.Equal(t, true, isSuccess)
	})

	s.T().Run("Delete | Invalid", func(t *testing.T) {
		s.classRepoMock.On("Delete", "").Return(false).Once()

		isSuccess := s.service.Delete("")

		assert.Equal(t, false, isSuccess)
	})
}

func (s *suiteClass) TestDeleteMany() {
	s.T().Run("DeleteMany | Valid", func(t *testing.T) {

		s.classRepoMock.On("DeleteMany", "1,2,3").Return(true).Once()

		isSuccess := s.service.DeleteMany("1,2,3")

		assert.Equal(t, true, isSuccess)
	})

	s.T().Run("DeleteMany | Invalid", func(t *testing.T) {
		s.classRepoMock.On("DeleteMany", "").Return(false).Once()

		isSuccess := s.service.DeleteMany("")

		assert.Equal(t, false, isSuccess)
	})
}

func TestSuiteClass(t *testing.T) {
	suite.Run(t, new(suiteClass))
}
