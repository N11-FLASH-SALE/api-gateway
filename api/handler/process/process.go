package process

import (
	"api/api/auth"
	pb "api/genproto/sale"
	"api/genproto/user"
	"api/models"

	"github.com/gin-gonic/gin"
)

// CreateProcess godoc
// @Security ApiKeyAuth
// @Description it will Create Process
// @Tags PROCESS
// @Param info body models.CreateProcessReq true "info"
// @Success 200 {object} user.GetCardsOfUserRes
// @Failure 401 {object} string "Invalid data"
// @Failure 401 {object} string "Invalid token"
// @Failure 401 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /process/buy [post]
func (h *newProcess) CreateProcess(c *gin.Context) {
	h.Log.Info("CreateProcess called")
	var req models.CreateProcessReq

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
	_, err = h.User.IsUserExist(c, &user.UserId{Id: userId})
	if err != nil {
		h.Log.Error("User not found", "error", err)
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	product, err := h.Product.GetProductById(c, &pb.ProductId{Id: req.ProductID})
	if err != nil {
		h.Log.Error("Product not found", "error", err)
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	cardInfo, err := h.Cards.GetCardAmount(c, &user.GetCardAmountReq{CardNumber: req.CardNumber})
	if err != nil {
		h.Log.Error("Error getting card amount", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if float64(req.Amount)*product.Price > cardInfo.Amount {
		h.Log.Error("Insufficient funds")
		c.JSON(400, gin.H{"error": "Insufficient funds"})
		return
	}

	_, err = h.Cards.UpdateCardAmount(c, &user.UpdateCardAmountReq{
		CardNumber: req.CardNumber,
		Amount:     cardInfo.Amount - product.Price*float64(req.Amount),
	})
	if err != nil {
		h.Log.Error("Error updating card amount", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = h.Bought.CreateBought(c, &pb.CreateBoughtRequest{
		UserId:        userId,
		ProductId:     req.ProductID,
		Amount:        req.Amount,
		CardNumber:    req.CardNumber,
		AmountOfMoney: product.Price * float64(req.Amount),
	})
	if err != nil {
		h.Log.Error("Error creating bought", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res1, err := h.Process.CreateProcess(c, &pb.CreateProcessRequest{
		UserId:    userId,
		ProductId: req.ProductID,
		Amount:    req.Amount,
	})

	if err != nil {
		h.Log.Error("Error creating process", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Process created successfully")
	c.JSON(200, res1)
}
