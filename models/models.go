package models

type CardRequest struct {
	CardNumber  string `json:"card_number" binding:"required"`
	ExpiresDate string `json:"expiration_date" binding:"required,datetime=01/06"`
	CCV         string `json:"ccv" binding:"required,len=3,numeric"`
}

type CreateProductRequest struct {
	Name              string   `json:"name" binding:"required"`
	Description       string   `json:"description" binding:"required"`
	PriceWithoutStock float64  `json:"price_without_stock" binding:"required"`
	Stock             int64    `json:"stock" binding:"required"`
	LimitOfProduct    int64    `json:"limit_of_product" binding:"required"`
	Size              []string `json:"size"`
	Color             []string `json:"color"`
	StartDate         string   `json:"start_date" binding:"required"`
	EndDate           string   `json:"end_date" binding:"required"`
}
