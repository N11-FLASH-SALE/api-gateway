package process

import (
	"api/api/auth"
	pb "api/genproto/sale"
	"api/genproto/user"
	"api/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateProcess godoc
// @Security ApiKeyAuth
// @Description it will Create Process
// @Tags PROCESS
// @Param info body models.CreateProcessReq true "info"
// @Success 200 {object} sale.ProcessResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 401 {object} string "Invalid token"
// @Failure 404 {object} string "Invalid data"
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
	_, err := h.Product.IsProductOk(c, &pb.ProductId{Id: req.ProductID})
	if err != nil {
		h.Log.Error("Product not found", "error", err)
		c.JSON(404, gin.H{"error": "Product not found"})
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

	if product.LimitOfProduct < req.Amount {
		h.Log.Error("Limit of product exceeded")
		c.JSON(400, gin.H{"error": "Limit of product exceeded"})
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
	_, err = h.Bought.CreateBought(c, &pb.CreateBoughtRequest{
		UserId:        userId,
		ProductId:     req.ProductID,
		Amount:        req.Amount,
		CardNumber:    req.CardNumber,
		AmountOfMoney: product.Price * float64(req.Amount),
		ProcessID:     res1.Id,
	})
	if err != nil {
		h.Log.Error("Error creating bought", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		h.Process.CancelProcess(c, &pb.CancelProcessRequest{Id: res1.Id})
		return
	}

	_, err = h.Product.UpdateLimitOfProduct(c, &pb.UpdateLimitOfProductRequest{Id: req.ProductID, LimitOfProduct: product.LimitOfProduct - req.Amount})
	if err != nil {
		h.Log.Error("Error updating product limit", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Notification.CreateNotification(c, &user.CreateNotificationsReq{UserId: userId, Message: "hello, you purchased product good luck!"})
	if err != nil {
		h.Log.Error("Error creating notification", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = h.Notification.CreateNotification(c, &user.CreateNotificationsReq{UserId: product.SellerId, Message: "hello, someone purchased your product good luck!"})
	if err != nil {
		h.Log.Error("Error creating notification", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Process created successfully")
	c.JSON(200, res1)
}

// GetProcessByProductId godoc
// @Security ApiKeyAuth
// @Description Get Process By Product Id
// @Tags PROCESS
// @Param product_id path string true "product_id"
// @Success 200 {object} sale.GetProcessByProductIdResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 404 {object} string "not found"
// @Failure 500 {object} string "Server error"
// @Router /process/products/{product_id} [get]
func (h *newProcess) GetProcessByProductId(c *gin.Context) {
	h.Log.Info("GetProcessByProductId called")
	id := c.Param("product_id")
	if len(id) == 0 {
		h.Log.Error("Invalid data")
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	_, err := h.Product.IsProductOk(c, &pb.ProductId{Id: id})
	if err != nil {
		h.Log.Error("Product not found", "error", err)
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	res, err := h.Process.GetProcessByProductId(c, &pb.GetProcessByProductIdRequest{ProductId: id})
	if err != nil {
		h.Log.Error("Error getting process", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Process found successfully")
	c.JSON(200, res)
}

// GetProcessOfUserByProductId godoc
// @Security ApiKeyAuth
// @Description Get Process Of User By Product Id
// @Tags PROCESS
// @Param product_id path string true "product_id"
// @Param user_id path string true "user_id"
// @Success 200 {object} sale.GetProcessOfUserByProductIdResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 404 {object} string "not found"
// @Failure 500 {object} string "Server error"
// @Router /process/user/{product_id}/{user_id} [get]
func (h *newProcess) GetProcessOfUserByProductId(c *gin.Context) {
	h.Log.Info("GetProcessOfUserByProductId called")
	id := c.Param("product_id")
	if len(id) == 0 {
		h.Log.Error("Invalid data")
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	_, err := h.Product.IsProductOk(c, &pb.ProductId{Id: id})
	if err != nil {
		h.Log.Error("Product not found", "error", err)
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	userId := c.Param("user_id")
	if len(userId) == 0 {
		h.Log.Error("Invalid data")
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	_, err = h.User.IsUserExist(c, &user.UserId{Id: userId})
	if err != nil {
		h.Log.Error("User not found", "error", err)
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	res, err := h.Process.GetProcessOfUserByProductId(c, &pb.GetProcessOfUserByProductIdRequest{
		ProductId: id,
		UserId:    userId,
	})
	if err != nil {
		h.Log.Error("Error getting process", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Process found successfully")
	c.JSON(200, res)
}

// GetProcessByUserId godoc
// @Security ApiKeyAuth
// @Description Get Process By User Id
// @Tags PROCESS
// @Success 200 {object} sale.GetProcessByUserIdResponse
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /process [get]
func (h *newProcess) GetProcessByUserId(c *gin.Context) {
	h.Log.Info("GetProcessByUserId called")

	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	res, err := h.Process.GetProcessByUserId(c, &pb.GetProcessByUserIdRequest{UserId: userId})
	if err != nil {
		h.Log.Error("Error getting process", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Process found successfully")
	c.JSON(200, res)
}

// GetProcessById godoc
// @Security ApiKeyAuth
// @Description Get Process By Id
// @Tags PROCESS
// @Param id path string true "id"
// @Success 200 {object} sale.GetProcessByIdResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /process/{id} [get]
func (h *newProcess) GetProcessById(c *gin.Context) {
	h.Log.Info("GetProcessById called")
	id := c.Param("id")
	if len(id) == 0 {
		h.Log.Error("Invalid data")
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	res, err := h.Process.GetProcessById(c, &pb.GetProcessByIdRequest{Id: id})
	if err != nil {
		h.Log.Error("Error getting process", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Process found successfully")
	c.JSON(200, res)
}

// UpdateProcess godoc
// @Security ApiKeyAuth
// @Description Update Process
// @Tags PROCESS
// @Param id path string true "id"
// @Param info body models.UpdateProcessReq true "info"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /process/{id} [put]
func (h *newProcess) UpdateProcess(c *gin.Context) {
	h.Log.Info("UpdateProcess called")
	var req models.UpdateProcessReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Invalid data", "error", err)
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	id := c.Param("id")
	if len(id) == 0 {
		h.Log.Error("Invalid data")
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	resfirst, err := h.Process.GetProcessById(c, &pb.GetProcessByIdRequest{Id: id})
	if resfirst.Status == "Cancelled" {
		h.Log.Error("Process is already Cancelled", "error", err)
		c.JSON(404, gin.H{"error": "Process is already Cancelled"})
		return
	}

	_, err = h.Process.UpdateProcess(c, &pb.UpdateProcessRequest{Id: id, Status: req.Status})
	if err != nil {
		h.Log.Error("Error updating process", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Notification.CreateNotification(c, &user.CreateNotificationsReq{UserId: resfirst.UserId, Message: fmt.Sprintf("your product is %s", req.Status)})
	if err != nil {
		h.Log.Error("Error creating notification", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Process updated successfully")
	c.JSON(200, gin.H{"message": "Process updated successfully"})
}

// CancelProcess godoc
// @Security ApiKeyAuth
// @Description Cancel Process
// @Tags PROCESS
// @Param id path string true "id"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /process/{id} [delete]
func (h *newProcess) CancelProcess(c *gin.Context) {
	h.Log.Info("CancelProcess called")
	id := c.Param("id")
	if len(id) == 0 {
		h.Log.Error("Invalid data")
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	resfirst, err := h.Process.GetProcessById(c, &pb.GetProcessByIdRequest{Id: id})
	if resfirst.Status != "Pending" {
		h.Log.Error("you can not cancel process", "error", err)
		c.JSON(404, gin.H{"error": "you can not cancel process"})
		_, err = h.Notification.CreateNotification(c, &user.CreateNotificationsReq{UserId: resfirst.UserId, Message: "you can not not cancel process because it is already above pending"})
		if err != nil {
			h.Log.Error("Error creating notification", "error", err)
			return
		}
		return
	}

	res, err := h.Process.CancelProcess(c, &pb.CancelProcessRequest{Id: id})
	if err != nil {
		h.Log.Error("Error canceling process", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res1, err := h.Bought.GetBoughtByProcessId(c, &pb.GetBoughtByProcessIdReq{ProcessId: id})
	if err != nil {
		h.Log.Error("Error getting bought", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res2, err := h.Product.GetProductById(c, &pb.ProductId{Id: res1.ProductId})
	if err != nil {
		h.Log.Error("Error getting product", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cardres, err := h.Cards.GetCardAmount(c, &user.GetCardAmountReq{CardNumber: res1.CardNumber})
	if err != nil {
		h.Log.Error("Error getting card amount", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Cards.UpdateCardAmount(c, &user.UpdateCardAmountReq{CardNumber: res1.CardNumber, Amount: res2.Price*float64(res.Amount) + cardres.Amount})
	if err != nil {
		h.Log.Error("Error updating card", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Notification.CreateNotification(c, &user.CreateNotificationsReq{UserId: resfirst.UserId, Message: "your purchase has canceled successfully"})
	if err != nil {
		h.Log.Error("Error creating notification", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Notification.CreateNotification(c, &user.CreateNotificationsReq{UserId: res2.SellerId, Message: "someone cancelled your products"})
	if err != nil {
		h.Log.Error("Error creating notification", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.Log.Info("Process canceled successfully")
	c.JSON(200, gin.H{"message": "Process canceled successfully"})
}
