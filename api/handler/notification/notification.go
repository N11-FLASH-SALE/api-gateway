package notification

import (
	"api/api/auth"
	"api/genproto/user"

	"github.com/gin-gonic/gin"
)

// GetAllNotifications godoc
// @Security ApiKeyAuth
// @Summary Get All Notifications
// @Description it will Get All Notifications
// @Tags NOTIFICATION
// @Success 200 {object} string "message"
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /notifications/all [get]
func (h *newNotifications) GetAllNotifications(c *gin.Context) {
	h.Log.Info("GetAllNotifications called")
	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	res, err := h.Notification.GetAllNotifications(c, &user.GetNotificationsReq{UserId: userId})
	if err != nil {
		h.Log.Error("Error getting all notifications", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("All notifications retrieved successfully")
	c.JSON(200, res)
}

// GetAndMarkNotificationAsRead godoc
// @Security ApiKeyAuth
// @Summary Get And Mark Notification As Read
// @Description it will Get And Mark Notification As Read
// @Tags NOTIFICATION
// @Success 200 {object} string "message"
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /notifications/unreaden [get]
func (h *newNotifications) GetAndMarkNotificationAsRead(c *gin.Context) {
	h.Log.Info("GetAndMarkNotificationAsRead called")
	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	res, err := h.Notification.GetAndMarkNotificationAsRead(c, &user.GetAndMarkNotificationAsReadReq{UserId: userId})
	if err != nil {
		h.Log.Error("Error getting and marking notification as read", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Notification retrieved and marked as read successfully")
	c.JSON(200, res)
}
