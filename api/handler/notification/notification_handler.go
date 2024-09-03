package notification

import (
	"api/genproto/user"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type newNotifications struct {
	Notification user.NotificationsClient
	Log          *slog.Logger
}

func NewNotificationsMethods(notification user.NotificationsClient, log *slog.Logger) NewNotification {
	return &newNotifications{
		Notification: notification,
		Log:          log,
	}
}

type NewNotification interface {
	GetAllNotifications(*gin.Context)
	GetAndMarkNotificationAsRead(*gin.Context)
}
