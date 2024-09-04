package cards

import (
	"api/api/auth"
	pb "api/genproto/user"
	"api/models"

	"github.com/gin-gonic/gin"
)

// CreateCards godoc
// @Security ApiKeyAuth
// @Summary Create Cards
// @Description it will Create Cards
// @Tags CARDS
// @Param info body models.CardRequest true "info"
// @Success 200 {object} user.CreateCardRes
// @Failure 400 {object} string "Invalid data"
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /cards [post]
func (h *newCards) CreateCards(c *gin.Context) {
	h.Log.Info("CreateCards called")
	req := models.CardRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Invalid request", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	res, err := h.Cards.CreateCard(c, &pb.CreateCardReq{UserId: userId, CardNumber: req.CardNumber, ExpirationDate: req.ExpiresDate, SecurityCode: req.CCV})
	if err != nil {
		h.Log.Error("Error creating card", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Card created successfully")
	c.JSON(200, res)
}

// GetUserCards godoc
// @Security ApiKeyAuth
// @Summary Get User Cards
// @Description it will Get User Cards
// @Tags CARDS
// @Success 200 {object} user.GetCardsOfUserRes
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /cards [get]
func (h *newCards) GetUserCards(c *gin.Context) {
	h.Log.Info("GetUserCards called")
	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	res, err := h.Cards.GetCardsOfUser(c, &pb.GetCardsOfUserReq{UserId: userId})
	if err != nil {
		h.Log.Error("Error getting user cards", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("User cards retrieved successfully")
	c.JSON(200, res)
}

// GetAmountOfUserCard godoc
// @Security ApiKeyAuth
// @Summary Get Amount Of User Card
// @Description it will Get Amount Of User Card
// @Tags CARDS
// @Param card_number path string true "card_number"
// @Success 200 {object} user.GetCardsOfUserRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /cards/amount/{card_number} [get]
func (h *newCards) GetAmountOfUserCard(c *gin.Context) {
	h.Log.Info("GetAmountOfUserCard called")

	num := c.Param("card_number")
	if len(num) == 0 {
		h.Log.Error("Invalid card number")
		c.JSON(400, gin.H{"error": "Invalid card number"})
		return
	}
	res, err := h.Cards.GetCardAmount(c, &pb.GetCardAmountReq{CardNumber: num})
	if err != nil {
		h.Log.Error("Error getting card amount", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Card amount retrieved successfully")
	c.JSON(200, res)
}

// DeleteCard godoc
// @Security ApiKeyAuth
// @Description it will Delete Card
// @Tags CARDS
// @Param card_number path string true "card_number"
// @Success 200 {object} string "message"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /cards/{card_number} [delete]
func (h *newCards) DeleteCard(c *gin.Context) {
	h.Log.Info("DeleteCard called")

	num := c.Param("card_number")
	if len(num) == 0 {
		h.Log.Error("Invalid card number")
		c.JSON(400, gin.H{"error": "Invalid card number"})
		return
	}
	_, err := h.Cards.DeleteCard(c, &pb.DeleteCardReq{CardNumber: num})
	if err != nil {
		h.Log.Error("Error deleting card", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Card deleted successfully")
	c.JSON(200, gin.H{"message": "Card deleted successfully"})
}
