package feedback

import (
	"api/api/auth"
	pb "api/genproto/sale"
	"api/genproto/user"
	"api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateFeedback godoc
// @Security ApiKeyAuth
// @Summary Create Feedback
// @Description it will Create Feedback
// @Tags FEEDBACK
// @Param product_id path string true "product_id"
// @Param info body models.CreateFeedback true "info"
// @Success 200 {object} sale.GetProductByIdResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 401 {object} string "Invalid token"
// @Failure 404 {object} string "not found"
// @Failure 500 {object} string "Server error"
// @Router /feedback/{product_id} [post]
func (h *newFeedbacks) CreateFeedback(c *gin.Context) {
	h.Log.Info("CreateFeedback called")
	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	prid := c.Param("product_id")
	if len(prid) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		h.Log.Error("Product ID is required")
		return
	}
	_, err = h.Product.IsProductOk(c, &pb.ProductId{Id: prid})
	if err != nil {
		h.Log.Error("Product not found", "error", err)
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	var req models.CreateFeedback
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Invalid request", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Feedback.CreateFeedback(c, &pb.CreateFeedbackRequest{
		UserId:      userId,
		ProductId:   prid,
		Rating:      req.Rating,
		Description: req.Description,
	})
	if err != nil {
		h.Log.Error("Error creating feedback", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = h.Notification.CreateNotification(c, &user.CreateNotificationsReq{UserId: userId, Message: fmt.Sprintf("you created a feedback for product id: %s", prid)})
	if err != nil {
		h.Log.Error("Error creating notification", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Feedback created successfully")
	c.JSON(200, res)
}

// GetFeedback godoc
// @Security ApiKeyAuth
// @Summary Get Feedback
// @Description it will Get Feedback
// @Tags FEEDBACK
// @Param product_id path string true "product_id"
// @Success 200 {object} sale.GetProductByIdResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 404 {object} string "not found"
// @Failure 500 {object} string "Server error"
// @Router /feedback/{product_id} [get]
func (h *newFeedbacks) GetFeedback(c *gin.Context) {
	h.Log.Info("GetFeedback called")
	prid := c.Param("product_id")
	if len(prid) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		h.Log.Error("Product ID is required")
		return
	}
	_, err := h.Product.IsProductOk(c, &pb.ProductId{Id: prid})
	if err != nil {
		h.Log.Error("Product not found", "error", err)
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	res, err := h.Feedback.GetFeedback(c, &pb.GetFeedbackRequest{
		ProductId: prid,
	})
	if err != nil {
		h.Log.Error("Error getting feedback", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Feedback retrieved successfully")
	c.JSON(200, res)
}

// GetFeedbackOfUser godoc
// @Security ApiKeyAuth
// @Summary Get Feedback Of User
// @Description it will Get Feedback Of User
// @Tags FEEDBACK
// @Success 200 {object} sale.GetProductByIdResponse
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /feedback [get]
func (h *newFeedbacks) GetFeedbackOfUser(c *gin.Context) {
	h.Log.Info("GetFeedbackOfUser called")
	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	res, err := h.Feedback.GetFeedbackOfUser(c, &pb.GetFeedbackOfUserRequest{UserId: userId})
	if err != nil {
		h.Log.Error("Error getting feedback", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Feedback retrieved successfully")
	c.JSON(200, res)
}
