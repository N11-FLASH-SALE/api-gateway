package models

// card
type CardRequest struct {
	CardNumber  string `json:"card_number" binding:"required"`
	ExpiresDate string `json:"expiration_date" binding:"required,datetime=01/06"`
	CCV         string `json:"ccv" binding:"required,len=3,numeric"`
}

// product
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

type GetProductReq struct {
	Name     string  `json:"name"`
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	Stock    int64   `json:"stock"`
	SellerID string  `json:"seller_id"`
}

type UpdateProductRequest struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Stock             int64    `json:"stock"`
	PriceWithoutStock float64  `json:"price_without_stock"`
	LimitOfProduct    int64    `json:"limit_of_product"`
	Size              []string `json:"size"`
	Color             []string `json:"color"`
	StartDate         string   `json:"start_date"`
	EndDate           string   `json:"end_date"`
}

type DeletePhoto struct {
	PhotoUrl string `json:"photo_url"`
}

// process
type CreateProcessReq struct {
	ProductID  string `json:"product_id" binding:"required"`
	Amount     int64  `json:"amount" binding:"required"`
	CardNumber string `json:"card_number" binding:"required"`
}

type UpdateProcessReq struct {
	Status string `json:"status" binding:"required"`
}

// wishlist

type WishList struct {
	ID                string   `json:"id"`
	ProductID         string   `json:"product_id"`
	ProductName       string   `json:"product_name"`
	Description       string   `json:"description"`
	Price             float64  `json:"price"`
	Stock             int64    `json:"stock"`
	PriceWithoutStock float64  `json:"price_without_stock"`
	LimitOfProduct    int64    `json:"limit_of_product"`
	Size              []string `json:"size"`
	Color             []string `json:"color"`
	StartDate         string   `json:"start_date"`
	EndDate           string   `json:"end_date"`
	SellerID          string   `json:"seller_id"`
	PhotoURL          []string `json:"photo_url"`
}

type WishListRes struct {
	UserID string      `json:"user_id"`
	Wishes []*WishList `json:"wish_list"`
}
