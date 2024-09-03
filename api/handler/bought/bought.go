package bought

import (
	"api/api/auth"
	pb "api/genproto/sale"

	"github.com/gin-gonic/gin"
)

// GetBought godoc
// @Security ApiKeyAuth
// @Summary Get Bought
// @Description it will Get Bought
// @Tags BOUGHT
// @Param product_id path string true "product_id"
// @Success 200 {object} string "message"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /bought/{product_id} [get]
func (h *newBoughts) GetBought(c *gin.Context) {
	h.Log.Info("GetBought called")
	productID := c.Param("product_id")
	if len(productID) == 0 {
		h.Log.Error("Invalid data")
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	res, err := h.Bought.GetBought(c, &pb.GetBoughtRequest{ProductId: productID})
	if err != nil {
		h.Log.Error("Error getting bought", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Bought retrieved successfully")
	c.JSON(200, res)
}

// GetBoughtOfUser godoc
// @Security ApiKeyAuth
// @Summary Get Bought Of User
// @Description it will Get Bought Of User
// @Tags BOUGHT
// @Success 200 {object} string "message"
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /bought [get]
func (h *newBoughts) GetBoughtOfUser(c *gin.Context) {
	h.Log.Info("GetBoughtOfUser called")
	token := c.GetHeader("Authorization")
	userID, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	res, err := h.Bought.GetBoughtOfUser(c, &pb.GetBoughtOfUserRequest{UserId: userID})
	if err != nil {
		h.Log.Error("Error getting bought of user", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Boughts of user retrieved successfully")
	c.JSON(200, res)
}
