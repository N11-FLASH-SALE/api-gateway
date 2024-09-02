package products

import (
	"api/api/auth"
	pb "api/genproto/sale"
	"api/models"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Security ApiKeyAuth
// @Summary Create Product
// @Description it will Create Product
// @Tags PRODUCTS
// @Param info body models.CreateProductRequest true "info"
// @Success 200 {object} sale.ProductId
// @Failure 400 {object} string "Invalid data"
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /products [post]
func (h *newProducts) CreateProduct(c *gin.Context) {
	h.Log.Info("CreateProduct called")
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Invalid request", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token := c.GetHeader("Authorization")
	sellerId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	res, err := h.Product.CreateProduct(c, &pb.CreateProductRequest{
		Name:              req.Name,
		Description:       req.Description,
		PriceWithoutStock: req.PriceWithoutStock,
		Stock:             req.Stock,
		LimitOfProduct:    req.LimitOfProduct,
		Size:              req.Size,
		Color:             req.Color,
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
		SellerId:          sellerId,
	})

	if err != nil {
		h.Log.Error("Error creating product", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Product created successfully")
	c.JSON(200, res)
}

func (h *newProducts) GetProductsList(*gin.Context) {

}

func (h *newProducts) GetProductByID(*gin.Context) {

}
