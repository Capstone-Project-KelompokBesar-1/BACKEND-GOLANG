package repositories

import (
	"ourgym/helpers"
	"ourgym/models"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewTransactionRepository(gormDB *gorm.DB, userRepo UserRepository, paymentMethodRepo PaymentMethodRepository) TransactionRepository {
	return &TransactionRepositoryImpl{
		db:                gormDB,
		userRepo:          userRepo,
		paymentMethodRepo: paymentMethodRepo,
	}
}

type TransactionRepositoryImpl struct {
	db                *gorm.DB
	userRepo          UserRepository
	paymentMethodRepo PaymentMethodRepository
}

func (tr *TransactionRepositoryImpl) GetAll() []models.Transaction {
	var transactions []models.Transaction

	tr.db.Preload(clause.Associations).Order("updated_at desc").Find(&transactions)

	return transactions
}

func (tr *TransactionRepositoryImpl) GetHistory() []models.Transaction {
	var transactions []models.Transaction

	tr.db.Preload(clause.Associations).Find(&transactions, "updated_at >= ? ORDER BY updated_at DESC", time.Now().Add(7*-24*time.Hour))

	return transactions
}

func (tr *TransactionRepositoryImpl) GetByUserID(userID, status string) []models.Transaction {
	var transactions []models.Transaction

	if status != "" {
		tr.db.Preload(clause.Associations).Find(&transactions, "user_id = ? && status = ?", userID, status)
	} else {
		tr.db.Preload(clause.Associations).Find(&transactions, "user_id = ?", userID)
	}

	return transactions
}

func (tr *TransactionRepositoryImpl) GetByID(id string) models.Transaction {
	var transaction models.Transaction

	tr.db.Preload(clause.Associations).First(&transaction, "id = ?", id)

	return transaction
}

func (tr *TransactionRepositoryImpl) Create(transactionRequest models.Transaction) (snap.Response, error) {
	transactionRequest.ID = uuid.NewString()
	transactionRequest.Status = "pending"

	user := tr.userRepo.GetOneByFilter("id", transactionRequest.UserID)
	paymentMethodIdString := strconv.FormatUint(uint64(transactionRequest.PaymentMethodID), 10)
	paymentMethod := tr.paymentMethodRepo.GetByID(paymentMethodIdString)

	s := helpers.InitMidtransSnap()

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionRequest.ID,
			GrossAmt: int64(transactionRequest.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			LName: ".",
			Email: user.Email,
			Phone: user.Phone,
		},
		EnabledPayments: []snap.SnapPaymentType{snap.SnapPaymentType(paymentMethod.Code)},
	}

	snapRes, err := s.CreateTransaction(req)

	if err != nil {
		return snap.Response{}, err
	}

	var transaction models.Transaction

	rec := tr.db.Create(&transactionRequest)

	rec.Last(&transaction)

	if rec.Error != nil {
		return snap.Response{}, rec.Error
	}

	return *snapRes, nil
}

func (tr *TransactionRepositoryImpl) UpdatedByMidtransAPI(transactionRequest models.Transaction) error {
	transaction := tr.GetByID(transactionRequest.ID)

	transaction.Status = transactionRequest.Status

	rec := tr.db.Save(&transaction)

	if rec.Error != nil {
		return rec.Error
	}

	return nil
}

func (tr *TransactionRepositoryImpl) Update(id string, transactionRequest models.Transaction) models.Transaction {
	transaction := tr.GetByID(id)

	transaction.User.ID = transactionRequest.UserID
	transaction.Class.ID = transactionRequest.ClassID
	transaction.PaymentMethod.ID = transactionRequest.PaymentMethodID
	transaction.Amount = transactionRequest.Amount
	transaction.Status = transactionRequest.Status

	rec := tr.db.Save(&transaction)

	rec.Preload(clause.Associations).Last(&transaction)

	return transaction
}

func (tr *TransactionRepositoryImpl) Delete(id string) bool {
	transaction := tr.GetByID(id)

	rec := tr.db.Delete(&transaction)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}

func (tr *TransactionRepositoryImpl) DeleteMany(ids string) bool {
	userIds := strings.Split(ids, ",")

	rec := tr.db.Delete(&models.Transaction{}, "id IN (?)", userIds)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}

func (tr *TransactionRepositoryImpl) CountTotalIncome() int64 {
	var totalIncome int64

	tr.db.Raw("SELECT sum(amount) FROM transactions WHERE (status = 'settlement' OR status = 'capture') AND updated_at >= ?", time.Now().Add(7*-24*time.Hour)).Scan(&totalIncome)

	return int64(totalIncome)
}
