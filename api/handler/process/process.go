package process

import (
	"api/api/auth"
	pb "api/genproto/sale"
	"api/models"

	"github.com/gin-gonic/gin"
)

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
