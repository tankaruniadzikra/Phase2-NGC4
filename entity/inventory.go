package entity

type Inventory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ItemCode    string `json:"itemCode"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
