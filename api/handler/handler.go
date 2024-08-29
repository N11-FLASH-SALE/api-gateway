package handler

import (
	"api/genproto/sale"
	"api/genproto/user"
	"log/slog"

	"github.com/casbin/casbin/v2"
)

type Handler struct {
	User         user.UserClient
	Product      sale.ProductClient
	Process      sale.ProcessClient
	Wishlis      sale.WishlistClient
	Feedback     sale.FeedbackClient
	Bought       sale.BoughtClient
	Notification user.NotificationsClient
	Log          *slog.Logger
	Enforcer     *casbin.Enforcer
}


