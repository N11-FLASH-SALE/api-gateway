package models

type CardRequest struct {
	CardNumber  string `json:"card_number" binding:"required"`
	ExpiresDate string `json:"expiration_date" binding:"required,datetime=01/06"`
	CCV         string `json:"ccv" binding:"required,len=3,numeric"`
}
