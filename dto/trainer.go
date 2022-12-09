package dto

type TrainerResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Photo       string `json:"photo"`
	Expertise   string `json:"expertise"`
	Description string `json:"description"`
}
