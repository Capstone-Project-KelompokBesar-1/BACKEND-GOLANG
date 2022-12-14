package services

import (
	"ourgym/dto"
	"ourgym/repositories"
)

func NewDashboardService(
	userRepo repositories.UserRepository,
	trainerRepo repositories.TrainerRepository,
	classRepo repositories.ClassRepository,
	transactionRepo repositories.TransactionRepository,
) DashboardService {
	return &DashboardServiceImpl{
		userRepo:        userRepo,
		trainerRepo:     trainerRepo,
		classRepo:       classRepo,
		transactionRepo: transactionRepo,
	}
}

type DashboardServiceImpl struct {
	userRepo        repositories.UserRepository
	trainerRepo     repositories.TrainerRepository
	classRepo       repositories.ClassRepository
	transactionRepo repositories.TransactionRepository
}

func (ds *DashboardServiceImpl) GetData() dto.DashboardResponse {
	totalUser := ds.userRepo.CountUser()
	totalTrainer := ds.trainerRepo.CountTrainer()
	totalClass := ds.classRepo.CountClass()
	totalIncome := ds.transactionRepo.CountTotalIncome()

	dashboardData := dto.DashboardResponse{
		TotalUser:    totalUser,
		TotalTrainer: totalTrainer,
		TotalClass:   totalClass,
		TotalIncome:  totalIncome,
	}

	return dashboardData
}
