package feedback

import (
	"api/genproto/sale"
	"api/genproto/user"
	"log/slog"
)

type newFeedbacks struct {
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

func NewFeedbacksMethods(
	cards user.CardsClient,
	user user.UserClient,
	product sale.ProductClient,
	process sale.ProcessClient,
	wishlist sale.WishlistClient,
	feedback sale.FeedbackClient,
	bought sale.BoughtClient,
	notification user.NotificationsClient,
	log *slog.Logger) NewFeedback {
	return &newFeedbacks{
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

type NewFeedback interface {
}
