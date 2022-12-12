package dto

type ClassResponse struct {
	ID           uint             `json:"id"`
	TrainerID    uint             `json:"trainer_id"`
	Trainer      TrainerResponse  `json:"trainer"`
	CategoryID   uint             `json:"category_id"`
	Category     CategoryResponse `json:"category"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	TotalMeeting int              `json:"total_meeting"`
	Thumbnail    string           `json:"thumbnail"`
	Type         string           `json:"type"`
	Price        int              `json:"price"`
}

type ClassForTransactionResponse struct {
	ID           uint   `json:"id"`
	TrainerID    uint   `json:"trainer_id"`
	CategoryID   uint   `json:"category_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	TotalMeeting int    `json:"total_meeting"`
	Thumbnail    string `json:"thumbnail"`
	Type         string `json:"type"`
	Price        int    `json:"price"`
}
