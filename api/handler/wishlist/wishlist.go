package wishlist

import (
	"api/api/auth"
	pb "api/genproto/sale"
	"api/genproto/user"
	"api/models"

	"github.com/gin-gonic/gin"
)

// CreateWishlist godoc
// @Security ApiKeyAuth
// @Description Create Wish list
// @Tags WISHLIST
// @Param product_id path string true "product_id"
// @Success 200 {object} sale.WishlistResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /wishlist/{product_id} [post]
func (h *newWishlists) CreateWishlist(c *gin.Context) {
	h.Log.Info("CreateWishlist called")
	id := c.Param("product_id")
	if len(id) == 0 {
		h.Log.Error("Invalid data")
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	res, err := h.Wishlist.CreateWishlist(c, &pb.CreateWishlistRequest{UserId: userId, ProductId: id})
	if err != nil {
		h.Log.Error("Error creating wishlist", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = h.Notification.CreateNotification(c, &user.CreateNotificationsReq{UserId: userId, Message: "you added product to your wishlist good luck to buy it:)"})
	if err != nil {
		h.Log.Error("Error creating notification", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Wishlist created successfully")
	c.JSON(200, res)
}

// GetWishlist godoc
// @Security ApiKeyAuth
// @Description Get Wish list
// @Tags WISHLIST
// @Success 200 {object} models.WishListRes
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /wishlist [get]
func (h *newWishlists) GetWishlist(c *gin.Context) {
	h.Log.Info("GetWishlist called")
	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	res, err := h.Wishlist.GetWishlist(c, &pb.GetWishlistRequest{UserId: userId})
	if err != nil {
		h.Log.Error("Error getting wishlist", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	req := models.WishListRes{
		UserID: userId,
	}
	for _, v := range res.Wishes {
		product, err := h.Product.GetProductById(c, &pb.ProductId{Id: v.ProductId})
		if err != nil {
			h.Log.Error("Error getting product", "error", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		req.Wishes = append(req.Wishes, &models.WishList{
			ID:                v.Id,
			ProductID:         product.Id,
			ProductName:       product.Name,
			Description:       product.Description,
			Price:             product.Price,
			Stock:             product.Stock,
			PriceWithoutStock: product.PriceWithoutStock,
			LimitOfProduct:    product.LimitOfProduct,
			Size:              product.Size,
			Color:             product.Color,
			StartDate:         product.StartDate,
			EndDate:           product.EndDate,
			SellerID:          product.SellerId,
			PhotoURL:          product.Photos,
		})
	}
	h.Log.Info("Wishlist retrieved successfully")
	c.JSON(200, req)
}

// GetWishlistById godoc
// @Security ApiKeyAuth
// @Description Get Wish list By Id
// @Tags WISHLIST
// @Param id path string true "id"
// @Success 200 {object} sale.WishlistResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /wishlist/{id} [get]
func (h *newWishlists) GetWishlistById(c *gin.Context) {
	h.Log.Info("GetWishlistById called")
	id := c.Param("id")
	if len(id) == 0 {
		h.Log.Error("Invalid data")
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	req, err := h.Wishlist.GetWishlistById(c, &pb.GetWishlistByIdRequest{Id: id})
	if err != nil {
		h.Log.Error("Error getting wishlist by ID", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	product, err := h.Product.GetProductById(c, &pb.ProductId{Id: req.ProductId})
	if err != nil {
		h.Log.Error("Error getting product", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res := models.WishList{
		ID:                id,
		ProductID:         product.Id,
		ProductName:       product.Name,
		Description:       product.Description,
		Price:             product.Price,
		Stock:             product.Stock,
		PriceWithoutStock: product.PriceWithoutStock,
		LimitOfProduct:    product.LimitOfProduct,
		Size:              product.Size,
		Color:             product.Color,
		StartDate:         product.StartDate,
		EndDate:           product.EndDate,
		SellerID:          product.SellerId,
		PhotoURL:          product.Photos,
	}
	h.Log.Info("Wishlist retrieved successfully")
	c.JSON(200, res)
}
