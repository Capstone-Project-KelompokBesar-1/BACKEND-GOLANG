package services

import (
	"ourgym/dto"
	"ourgym/models"
	"ourgym/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type suiteUser struct {
	suite.Suite
	service      UserService
	userRepoMock *mocks.UserRepository
	userRequest  dto.UserRequest
	userModel    models.User
}

func (s *suiteUser) SetupSuite() {
	s.userRepoMock = &mocks.UserRepository{}

	s.service = NewUserService(s.userRepoMock)

	s.userRequest = dto.UserRequest{
		Name:     "testing user",
		Phone:    "0890980090",
		Email:    "testing@mail.com",
		Password: "1234",
	}

	s.userModel = models.User{
		ID:       1,
		Name:     "testing user",
		Phone:    "0890980090",
		Email:    "testing@mail.com",
		Password: "1234",
		IsAdmin:  false,
	}

}

func (s *suiteUser) TestGetAll() {
	s.T().Run("GetAll | Valid", func(t *testing.T) {
		s.userRepoMock.On("GetAll", "").Return([]models.User{s.userModel}).Once()

		users := s.service.GetAll("")

		assert.Equal(t, 1, len(users))
	})

	s.T().Run("GetAll | Invalid", func(t *testing.T) {
		s.userRepoMock.On("GetAll", "").Return([]models.User{}).Once()

		users := s.service.GetAll("")

		assert.Equal(t, 0, len(users))
	})
}

func (s *suiteUser) TestGetByID() {
	s.T().Run("GetByID | Valid", func(t *testing.T) {
		s.userRepoMock.On("GetOneByFilter", "id", "1").Return(s.userModel).Once()

		user := s.service.GetByID("1")

		assert.Equal(t, uint(1), user.ID)
	})

	s.T().Run("GetByID | Invalid", func(t *testing.T) {
		s.userRepoMock.On("GetOneByFilter", "id", "").Return(models.User{}).Once()

		user := s.service.GetByID("")

		assert.Equal(t, uint(0), user.ID)
	})
}

func (s *suiteUser) TestCreate() {
	s.T().Run("Create | Invalid", func(t *testing.T) {
		s.userRepoMock.On("GetOneByFilter", "email", s.userRequest.Email).Return(s.userModel).Once()

		s.userRepoMock.On("Create", models.User{}).Return(models.User{}).Once()

		user, err := s.service.Create(s.userRequest)

		assert.Error(t, err)

		assert.Equal(t, uint(0), user.ID)
	})
}

func (s *suiteUser) TestUpdate() {
	s.T().Run("Update | Valid", func(t *testing.T) {
		userReq := models.FromUserRequestToUserModel(s.userRequest)

		s.userRepoMock.On("Update", "1", userReq).Return(s.userModel).Once()

		user := s.service.Update("1", s.userRequest)

		assert.Equal(t, uint(1), user.ID)
	})

	s.T().Run("Update | Invalid", func(t *testing.T) {
		s.userRepoMock.On("Update", "0", models.User{}).Return(models.User{}).Once()

		user := s.service.Update("0", dto.UserRequest{})

		assert.Equal(t, uint(0), user.ID)
	})
}

func (s *suiteUser) TestDelete() {
	s.T().Run("Delete | Valid", func(t *testing.T) {

		s.userRepoMock.On("Delete", "1").Return(true).Once()

		isSuccess := s.service.Delete("1")

		assert.Equal(t, true, isSuccess)
	})

	s.T().Run("Delete | Invalid", func(t *testing.T) {
		s.userRepoMock.On("Delete", "").Return(false).Once()

		isSuccess := s.service.Delete("")

		assert.Equal(t, false, isSuccess)
	})
}

func (s *suiteUser) TestDeleteMany() {
	s.T().Run("DeleteMany | Valid", func(t *testing.T) {

		s.userRepoMock.On("DeleteMany", "1,2,3").Return(true).Once()

		isSuccess := s.service.DeleteMany("1,2,3")

		assert.Equal(t, true, isSuccess)
	})

	s.T().Run("DeleteMany | Invalid", func(t *testing.T) {
		s.userRepoMock.On("DeleteMany", "").Return(false).Once()

		isSuccess := s.service.DeleteMany("")

		assert.Equal(t, false, isSuccess)
	})
}

func TestSuiteUser(t *testing.T) {
	suite.Run(t, new(suiteUser))
}
