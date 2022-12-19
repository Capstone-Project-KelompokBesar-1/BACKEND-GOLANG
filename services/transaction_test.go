package services

import (
	"ourgym/dto"
	"ourgym/models"
	"ourgym/repositories/mocks"
	"testing"

	"github.com/midtrans/midtrans-go/snap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type suiteTransaction struct {
	suite.Suite
	service             TransactionService
	transactionRepoMock *mocks.TransactionRepository
	transactionRequest  dto.TransactionRequest
	transactionModel    models.Transaction
	midtransRequest     dto.MidtransTransactionRequest
}

func (s *suiteTransaction) SetupSuite() {
	s.transactionRepoMock = &mocks.TransactionRepository{}

	s.service = NewTransactionService(s.transactionRepoMock)

	s.transactionRequest = dto.TransactionRequest{
		UserID:          1,
		ClassID:         1,
		PaymentMethodID: 1,
		Amount:          100000,
	}

	s.transactionModel = models.Transaction{
		ID:              "uheafeo2hfew8-92hhfa98ias-7h3f0ao1uaefr",
		UserID:          1,
		ClassID:         1,
		PaymentMethodID: 1,
		Amount:          100000,
		Status:          "pending",
	}

	s.midtransRequest = dto.MidtransTransactionRequest{
		TransactionStatus: "settlement",
		OrderID:           "adsfiojeaw-i0gj3waj0aisd-834jraojgfagh",
	}

}

func (s *suiteTransaction) TestGetAll() {
	s.T().Run("GetAll | Valid", func(t *testing.T) {
		s.transactionRepoMock.On("GetAll").Return([]models.Transaction{s.transactionModel}).Once()

		transactions := s.service.GetAll()

		assert.Equal(t, 1, len(transactions))
	})

	s.T().Run("GetAll | Invalid", func(t *testing.T) {
		s.transactionRepoMock.On("GetAll").Return([]models.Transaction{}).Once()

		transactions := s.service.GetAll()

		assert.Equal(t, 0, len(transactions))
	})
}

func (s *suiteTransaction) TestGetHistory() {
	s.T().Run("GetHistory | Valid", func(t *testing.T) {
		s.transactionRepoMock.On("GetHistory").Return([]models.Transaction{s.transactionModel}).Once()

		transactions := s.service.GetHistory()

		assert.Equal(t, 1, len(transactions))
	})

	s.T().Run("GetHistory | Invalid", func(t *testing.T) {
		s.transactionRepoMock.On("GetHistory").Return([]models.Transaction{}).Once()

		transactions := s.service.GetHistory()

		assert.Equal(t, 0, len(transactions))
	})
}

func (s *suiteTransaction) TestGetByUserID() {
	s.T().Run("GetByUserID | Valid", func(t *testing.T) {
		s.transactionRepoMock.On("GetByUserID", "1", "").Return([]models.Transaction{s.transactionModel}).Once()

		transactions := s.service.GetByUserID("1", "")

		assert.Equal(t, 1, len(transactions))
	})

	s.T().Run("GetByUserID | Invalid", func(t *testing.T) {
		s.transactionRepoMock.On("GetByUserID", "", "").Return([]models.Transaction{}).Once()

		transactions := s.service.GetByUserID("", "")

		assert.Equal(t, 0, len(transactions))
	})
}

func (s *suiteTransaction) TestGetByID() {
	s.T().Run("GetByID | Valid", func(t *testing.T) {
		s.transactionRepoMock.On("GetByID", "1").Return(s.transactionModel).Once()

		transaction := s.service.GetByID("1")

		assert.NotEmpty(t, transaction.ID)
	})

	s.T().Run("GetByID | Invalid", func(t *testing.T) {
		s.transactionRepoMock.On("GetByID", "").Return(models.Transaction{}).Once()

		transaction := s.service.GetByID("")

		assert.Empty(t, transaction.ID)
	})
}

func (s *suiteTransaction) TestCreate() {
	s.T().Run("Create | Valid", func(t *testing.T) {
		transactionReq := models.FromTransactionRequest(s.transactionRequest)

		snapResponse := snap.Response{
			Token:       "iejfisgjawejgaewfkds.asjdfoijewajfag3.9j392fhj9fadshf",
			RedirectURL: "https://testing.com",
		}

		s.transactionRepoMock.On("Create", transactionReq).Return(snapResponse, nil).Once()

		snap, err := s.service.Create(s.transactionRequest)

		assert.NoError(t, err)

		assert.NotEmpty(t, snap.RedirectURL)
	})
}

func (s *suiteTransaction) TestUpdate() {
	s.T().Run("Update | Valid", func(t *testing.T) {
		transactionReq := models.FromTransactionRequest(s.transactionRequest)

		s.transactionRepoMock.On("Update", "1", transactionReq).Return(s.transactionModel).Once()

		transaction := s.service.Update("1", s.transactionRequest)

		assert.NotEmpty(t, transaction.ID)
	})

	s.T().Run("Update | Invalid", func(t *testing.T) {
		s.transactionRepoMock.On("Update", "0", models.Transaction{}).Return(models.Transaction{}).Once()

		transaction := s.service.Update("0", dto.TransactionRequest{})

		assert.Empty(t, transaction.ID)
	})
}

func (s *suiteTransaction) TestDelete() {
	s.T().Run("Delete | Valid", func(t *testing.T) {

		s.transactionRepoMock.On("Delete", "1").Return(true).Once()

		isSuccess := s.service.Delete("1")

		assert.Equal(t, true, isSuccess)
	})

	s.T().Run("Delete | Invalid", func(t *testing.T) {
		s.transactionRepoMock.On("Delete", "").Return(false).Once()

		isSuccess := s.service.Delete("")

		assert.Equal(t, false, isSuccess)
	})
}

func (s *suiteTransaction) TestDeleteMany() {
	s.T().Run("DeleteMany | Valid", func(t *testing.T) {

		s.transactionRepoMock.On("DeleteMany", "1,2,3").Return(true).Once()

		isSuccess := s.service.DeleteMany("1,2,3")

		assert.Equal(t, true, isSuccess)
	})

	s.T().Run("DeleteMany | Invalid", func(t *testing.T) {
		s.transactionRepoMock.On("DeleteMany", "").Return(false).Once()

		isSuccess := s.service.DeleteMany("")

		assert.Equal(t, false, isSuccess)
	})
}

func TestSuiteTransaction(t *testing.T) {
	suite.Run(t, new(suiteTransaction))
}
