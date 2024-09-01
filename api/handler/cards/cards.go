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
	c.JSON(200, res)
}

func (h *newCards) GetUserCards(c *gin.Context) {

}

func (h *newCards) GetAmountOfUserCard(c *gin.Context) {

}
