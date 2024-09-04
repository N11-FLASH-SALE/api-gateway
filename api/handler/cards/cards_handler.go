package cards

import (
	"api/genproto/sale"
	"api/genproto/user"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type newCards struct {
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

func NewCardsMethods(
	cards user.CardsClient,
	user user.UserClient,
	product sale.ProductClient,
	process sale.ProcessClient,
	wishlist sale.WishlistClient,
	feedback sale.FeedbackClient,
	bought sale.BoughtClient,
	notification user.NotificationsClient,
	log *slog.Logger) NewCard {
	return &newCards{
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

type NewCard interface {
	CreateCards(*gin.Context)
	GetUserCards(*gin.Context)
	GetAmountOfUserCard(*gin.Context)
	DeleteCard(*gin.Context)
}
