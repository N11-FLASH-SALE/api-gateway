package handler

import (
	"api/api/handler/bought"
	"api/api/handler/cards"
	"api/api/handler/feedback"
	"api/api/handler/notification"
	"api/api/handler/process"
	"api/api/handler/products"
	"api/api/handler/wishlist"
	"api/api/middleware"
	"api/genproto/sale"
	"api/genproto/user"
	"log/slog"

	"github.com/casbin/casbin/v2"
)

type HandlerInterface interface {
	ProductMethods() products.NewProduct
	BoughtMethods() bought.NewBought
	CardsMethods() cards.NewCard
	FeedbackMethods() feedback.NewFeedback
	ProcessMethods() process.NewProcess
	WishlistMethods() wishlist.NewWishlist
	NotificationMethods() notification.NewNotification
	EnforcerMethods() middleware.CasbinPermission
}

type Handler struct {
	User         user.UserClient
	Product      sale.ProductClient
	Process      sale.ProcessClient
	Wishlist     sale.WishlistClient
	Feedback     sale.FeedbackClient
	Bought       sale.BoughtClient
	Notification user.NotificationsClient
	Cards        user.CardsClient
	Log          *slog.Logger
	Enforcer     *casbin.Enforcer
}

func (h *Handler) ProductMethods() products.NewProduct {
	return products.NewProductsMethods(
		h.Cards,
		h.User,
		h.Product,
		h.Process,
		h.Wishlist,
		h.Feedback,
		h.Bought,
		h.Notification,
		h.Log,
	)
}

func (h *Handler) BoughtMethods() bought.NewBought {
	return bought.NewBoughtsMethods(
		h.Cards,
		h.User,
		h.Product,
		h.Process,
		h.Wishlist,
		h.Feedback,
		h.Bought,
		h.Notification,
		h.Log,
	)
}

func (h *Handler) CardsMethods() cards.NewCard {
	return cards.NewCardsMethods(
		h.Cards,
		h.User,
		h.Product,
		h.Process,
		h.Wishlist,
		h.Feedback,
		h.Bought,
		h.Notification,
		h.Log,
	)
}

func (h *Handler) FeedbackMethods() feedback.NewFeedback {
	return feedback.NewFeedbacksMethods(
		h.Cards,
		h.User,
		h.Product,
		h.Process,
		h.Wishlist,
		h.Feedback,
		h.Bought,
		h.Notification,
		h.Log,
	)
}

func (h *Handler) ProcessMethods() process.NewProcess {
	return process.NewProcessMethods(
		h.Cards,
		h.User,
		h.Product,
		h.Process,
		h.Wishlist,
		h.Feedback,
		h.Bought,
		h.Notification,
		h.Log,
	)
}

func (h *Handler) WishlistMethods() wishlist.NewWishlist {
	return wishlist.NewWishlistsMethods(
		h.Cards,
		h.User,
		h.Product,
		h.Process,
		h.Wishlist,
		h.Feedback,
		h.Bought,
		h.Notification,
		h.Log,
	)
}

func (h *Handler) NotificationMethods() notification.NewNotification {
	return notification.NewNotificationsMethods(
		h.Notification,
		h.Log,
	)
}

func (h *Handler) EnforcerMethods() middleware.CasbinPermission {
	return middleware.NewCasbinPermission(h.Enforcer)
}
