package dto

type DTOClass struct {
	ID          uint        `json:"id"`
	TrainerID   uint        `json:"trainer_id"`
	Trainer     DTOTrainer  `json:"trainer"`
	CategoryID  uint        `json:"category_id"`
	Category    DTOCategory `json:"category"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Thumbnail   string      `json:"thumbnail"`
	Type        string      `json:"type"`
	Price       int         `json:"price"`
}
