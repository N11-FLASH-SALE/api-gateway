package wishlist

import (
	"api/genproto/sale"
	"api/genproto/user"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type newWishlists struct {
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

func NewWishlistsMethods(
	cards user.CardsClient,
	user user.UserClient,
	product sale.ProductClient,
	process sale.ProcessClient,
	wishlist sale.WishlistClient,
	feedback sale.FeedbackClient,
	bought sale.BoughtClient,
	notification user.NotificationsClient,
	log *slog.Logger) NewWishlist {
	return &newWishlists{
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

type NewWishlist interface {
	CreateWishlist(*gin.Context)
	GetWishlist(*gin.Context)
	GetWishlistById(*gin.Context)
}
