package services

import (
	"ourgym/dto"
	"ourgym/models"
	"ourgym/repositories"

	"github.com/midtrans/midtrans-go/snap"
)

func NewTransactionService(tr repositories.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		transactionRepository: tr,
	}
}

type TransactionServiceImpl struct {
	transactionRepository repositories.TransactionRepository
}

func (cs *TransactionServiceImpl) GetAll() []dto.TransactionResponse {
	transactions := cs.transactionRepository.GetAll()

	var transactionsResponse []dto.TransactionResponse

	for _, transaction := range transactions {
		transactionsResponse = append(transactionsResponse, transaction.ConvertToDTO())
	}

	return transactionsResponse
}

func (cs *TransactionServiceImpl) GetHistory() []dto.TransactionResponse {
	transactions := cs.transactionRepository.GetHistory()

	var transactionsResponse []dto.TransactionResponse

	for _, transaction := range transactions {
		transactionsResponse = append(transactionsResponse, transaction.ConvertToDTO())
	}

	return transactionsResponse
}

func (cs *TransactionServiceImpl) GetByUserID(userID, status string) []dto.TransactionResponse {
	transactions := cs.transactionRepository.GetByUserID(userID, status)

	var transactionsResponse []dto.TransactionResponse

	for _, transaction := range transactions {
		transactionsResponse = append(transactionsResponse, transaction.ConvertToDTO())
	}

	return transactionsResponse
}

func (cs *TransactionServiceImpl) GetByID(id string) dto.TransactionResponse {
	transaction := cs.transactionRepository.GetByID(id)
	return transaction.ConvertToDTO()
}

func (cs *TransactionServiceImpl) Create(transactionRequest dto.TransactionRequest) (snap.Response, error) {
	transactionModel := models.FromTransactionRequest(transactionRequest)

	snapRes, err := cs.transactionRepository.Create(transactionModel)

	return snapRes, err
}

func (cs *TransactionServiceImpl) UpdatedByMidtransAPI(midtransTransaction dto.MidtransTransactionRequest) error {
	transactionModel := models.Transaction{
		ID:     midtransTransaction.OrderID,
		Status: midtransTransaction.TransactionStatus,
	}

	err := cs.transactionRepository.UpdatedByMidtransAPI(transactionModel)

	return err
}

func (cs *TransactionServiceImpl) Update(id string, transactionRequest dto.TransactionRequest) dto.TransactionResponse {
	transactionModel := models.FromTransactionRequest(transactionRequest)

	transaction := cs.transactionRepository.Update(id, transactionModel)

	return transaction.ConvertToDTO()
}

func (cs *TransactionServiceImpl) Delete(id string) bool {
	return cs.transactionRepository.Delete(id)
}

func (cs *TransactionServiceImpl) DeleteMany(ids string) bool {
	return cs.transactionRepository.DeleteMany(ids)
}
