package dto

type DashboardResponse struct {
	TotalUser    int64 `json:"total_user"`
	TotalTrainer int64 `json:"total_trainer"`
	TotalClass   int64 `json:"total_class"`
	TotalIncome  int64 `json:"total_income"`
}
