package process

import (
	"api/genproto/sale"
	"api/genproto/user"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type newProcess struct {
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

func NewProcessMethods(
	cards user.CardsClient,
	user user.UserClient,
	product sale.ProductClient,
	process sale.ProcessClient,
	wishlist sale.WishlistClient,
	feedback sale.FeedbackClient,
	bought sale.BoughtClient,
	notification user.NotificationsClient,
	log *slog.Logger) NewProcess {
	return &newProcess{
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

type NewProcess interface {
	CreateProcess(*gin.Context)
	GetProcessByProductId(*gin.Context)
	GetProcessOfUserByProductId(*gin.Context)
	GetProcessByUserId(*gin.Context)
	GetProcessById(*gin.Context)
	UpdateProcess(*gin.Context)
	CancelProcess(*gin.Context)
}
