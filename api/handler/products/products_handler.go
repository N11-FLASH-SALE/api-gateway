package products

import (
	"api/genproto/sale"
	"api/genproto/user"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type newProducts struct {
	Cards        user.CardsClient
	User         user.UserClient
	Product      sale.ProductClient
	Process      sale.ProcessClient
	Wishlist     sale.WishlistClient
	Feedback     sale.FeedbackClient
	Bought       sale.BoughtClient
	Notification user.NotificationsClient
	Log          *slog.Logger
}

func NewProductsMethods(
	cards user.CardsClient,
	user user.UserClient,
	product sale.ProductClient,
	process sale.ProcessClient,
	wishlist sale.WishlistClient,
	feedback sale.FeedbackClient,
	bought sale.BoughtClient,
	notification user.NotificationsClient,
	log *slog.Logger) NewProduct {
	return &newProducts{
		Cards:        cards,
		User:         user,
		Product:      product,
		Process:      process,
		Wishlist:     wishlist,
		Feedback:     feedback,
		Bought:       bought,
		Notification: notification,
		Log:          log,
	}
}

type NewProduct interface {
	CreateProduct(*gin.Context)
	GetProductsList(*gin.Context)
	GetProductByID(*gin.Context)
	GetSellerProducts(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
	UploadProductPhoto(*gin.Context)
	DeleteProductPhoto(*gin.Context)
}
