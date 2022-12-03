package dto

type DTOTrainer struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Photo       string `json:"photo"`
	Expertise   string `json:"expertise"`
	Description string `json:"description"`
}
