package products

import (
	"api/api/auth"
	"api/config"
	pb "api/genproto/sale"
	"api/genproto/user"
	"api/models"
	"api/queue/kafka/producer"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

	_, err = h.User.IsUserExist(c, &user.UserId{Id: sellerId})
	if err != nil {
		h.Log.Error("Seller not found", "error", err)
		c.JSON(404, gin.H{"error": "Seller not found"})
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

// GetProductsList godoc
// @Security ApiKeyAuth
// @Summary Get Products List
// @Description it will Get Products List
// @Tags PRODUCTS
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Param info body models.GetProductReq true "info"
// @Success 200 {object} sale.GetProductResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 404 {object} string "Invalid user"
// @Failure 500 {object} string "Server error"
// @Router /products/list [post]
func (h *newProducts) GetProductsList(c *gin.Context) {
	h.Log.Info("GetProductsList called")
	var req models.GetProductReq

	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Invalid request", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(req.SellerID) > 0 {
		_, err := h.User.IsUserExist(c, &user.UserId{Id: req.SellerID})
		if err != nil {
			h.Log.Error("Seller not found", "error", err)
			c.JSON(404, gin.H{"error": "Seller not found"})
			return
		}
	}
	reqMain := pb.GetProductRequest{
		Name:     req.Name,
		MinPrice: req.MinPrice,
		MaxPrice: req.MaxPrice,
		Stock:    req.Stock,
		SellerId: req.SellerID,
	}

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		reqMain.Limit = int64(limit)
	} else {
		reqMain.Limit = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		reqMain.Offset = int64(offset)
	} else {
		reqMain.Offset = 0
	}

	res, err := h.Product.GetProduct(c, &reqMain)
	if err != nil {
		h.Log.Error("Error getting products list", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Products list retrieved successfully")
	c.JSON(200, res)
}

// GetProductByID godoc
// @Security ApiKeyAuth
// @Summary Get Product By ID
// @Description it will Get Product By ID
// @Tags PRODUCTS
// @Param id path string true "id"
// @Success 200 {object} sale.GetProductByIdResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /products/{id} [get]
func (h *newProducts) GetProductByID(c *gin.Context) {
	h.Log.Info("GetProductByID called")
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		h.Log.Error("Product ID is required")
		return
	}
	res, err := h.Product.GetProductById(c, &pb.ProductId{Id: id})
	if err != nil {
		h.Log.Error("Error getting product by ID", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Product retrieved successfully")
	c.JSON(200, res)
}

// GetSellerProducts godoc
// @Security ApiKeyAuth
// @Summary Get Seller Products
// @Description it will Get Seller Products
// @Tags PRODUCTS
// @Success 200 {object} sale.GetProductsByUserIdResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /products [get]
func (h *newProducts) GetSellerProducts(c *gin.Context) {
	h.Log.Info("GetSellerProducts called")
	token := c.GetHeader("Authorization")
	sellerId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	if len(sellerId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Seller ID is required"})
		h.Log.Error("Seller ID is required")
		return
	}
	res, err := h.Product.GetProductsByUserId(c, &pb.GetProductsByUserIdRequest{SellerId: sellerId})
	if err != nil {
		h.Log.Error("Error getting seller products", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Seller products retrieved successfully")
	c.JSON(200, res)
}

// UpdateProduct godoc
// @Security ApiKeyAuth
// @Summary Update Product
// @Description it will Update Product
// @Tags PRODUCTS
// @Param id path string true "id"
// @Param info body models.UpdateProductRequest true "info"
// @Success 200 {object} string "message"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /products/{id} [put]
func (h *newProducts) UpdateProduct(c *gin.Context) {
	h.Log.Info("UpdateProduct called")
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		h.Log.Error("Product ID is required")
		return
	}
	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Invalid request", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Product.GetProductById(c, &pb.ProductId{Id: id})
	if err != nil {
		h.Log.Error("Error getting product by ID", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	reqMain := pb.UpdateProductRequest{
		Id:                id,
		Name:              req.Name,
		Description:       req.Description,
		Stock:             req.Stock,
		PriceWithoutStock: req.PriceWithoutStock,
		Size:              req.Size,
		Color:             req.Color,
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
	}

	if req.Stock != 0 {
		if req.PriceWithoutStock != 0 {
			stockPercentage := float64(req.Stock) / 100
			reqMain.Price = req.PriceWithoutStock - req.PriceWithoutStock*stockPercentage
		} else {
			stockPercentage := float64(req.Stock) / 100
			reqMain.Price = res.PriceWithoutStock - res.PriceWithoutStock*stockPercentage
		}
	} else if req.PriceWithoutStock != 0 {
		stockPercentage := float64(res.Stock) / 100
		reqMain.Price = req.PriceWithoutStock - req.PriceWithoutStock*stockPercentage
	}

	// kafka
	writerKafka, err := producer.NewKafkaProducerInit([]string{"kafka:9092"})
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, err.Error())
		return
	}
	defer writerKafka.Close()
	msgBytes, err := json.Marshal(&reqMain)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, err.Error())
		return
	}
	err = writerKafka.Producermessage("update_product", msgBytes)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, err.Error())
		return
	}
	h.Log.Info("Product updated successfully")
	c.JSON(200, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct godoc
// @Security ApiKeyAuth
// @Summary Delete Product
// @Description it will Delete Product
// @Tags PRODUCTS
// @Param id path string true "id"
// @Success 200 {object} string "message"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /products/{id} [delete]
func (h *newProducts) DeleteProduct(c *gin.Context) {
	h.Log.Info("DeleteProduct called")
	var req pb.ProductId
	req.Id = c.Param("id")
	if len(req.Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		h.Log.Error("Product ID is required")
		return
	}
	writerKafka, err := producer.NewKafkaProducerInit([]string{"kafka:9092"})
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, err.Error())
		return
	}
	defer writerKafka.Close()
	msgBytes, err := json.Marshal(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, err.Error())
		return
	}
	err = writerKafka.Producermessage("delete_product", msgBytes)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, err.Error())
		return
	}
	h.Log.Info("Product deleted successfully")
	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}

// @Summary UploadProductPhoto
// @Security ApiKeyAuth
// @Description Upload Product Photo
// @Tags PRODUCTS
// @Accept multipart/form-data
// @Param product_id path string true "product_id"
// @Param file formData file true "UploadMediaForm"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /products/photo/{product_id} [post]
func (h *newProducts) UploadProductPhoto(c *gin.Context) {
	h.Log.Info("UploadProductPhoto called")

	Id := c.Param("product_id")
	if len(Id) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file"})
		return
	}
	defer file.Close()

	// minio start

	fileExt := filepath.Ext(header.Filename)
	println("\n File Ext:", fileExt)

	newFile := uuid.NewString() + fileExt
	minioClient, err := minio.New(config.Load().MINIO_URL, &minio.Options{
		Creds:  credentials.NewStaticV4("test", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	info, err := minioClient.PutObject(context.Background(), "products", newFile, file, header.Size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		c.AbortWithError(500, err)
		fmt.Println(err.Error())
		return
	}

	policy := fmt.Sprintf(`{
	 "Version": "2012-10-17",
	 "Statement": [
	  {
	   "Effect": "Allow",
	   "Principal": {
		"AWS": ["*"]
	   },
	   "Action": ["s3:GetObject"],
	   "Resource": ["arn:aws:s3:::%s/*"]
	  }
	 ]
	}`, "products")

	err = minioClient.SetBucketPolicy(context.Background(), "products", policy)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	madeUrl := fmt.Sprintf("http://%s/products/%s", config.Load().MINIO_URL, newFile)

	println("\n Info Bucket:", info.Bucket)

	// minio end

	req := pb.AddPhotosRequest{
		ProductId: Id,
		PhotoUrl:  madeUrl,
	}
	_, err = h.Product.AddPhotosToProduct(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}
	h.Log.Info("UploadMediaProduct finished successfully")
	c.JSON(200, gin.H{
		"minio url": madeUrl,
	})

}

// @Summary DeleteProductPhoto
// @Security ApiKeyAuth
// @Description Delete Product Photo
// @Tags PRODUCTS
// @Param product_id path string true "product_id"
// @Param url query string false "url"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /products/photo/{product_id} [delete]
func (h *newProducts) DeleteProductPhoto(c *gin.Context) {
	h.Log.Info("DeleteProductPhoto called")
	Id := c.Param("product_id")
	if len(Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		h.Log.Error("Product ID is required")
		return
	}
	var ph models.DeletePhoto
	ph.PhotoUrl = c.Query("url")

	prefix := fmt.Sprintf("http://%s/products/", config.Load().MINIO_URL)
	bucketName := "products"
	objectName := strings.TrimPrefix(ph.PhotoUrl, prefix)

	minioClient, err := minio.New(config.Load().MINIO_URL, &minio.Options{
		Creds:  credentials.NewStaticV4("test", "minioadmin", ""),
		Secure: false, // Set to true if using HTTPS
	})
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to minio"})
	}
	err = minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting photo"})
		return
	}
	req := pb.DeletePhotosRequest{
		ProductId: Id,
		PhotoUrl:  ph.PhotoUrl,
	}
	_, err = h.Product.DeletePhotosFromProduct(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}
	h.Log.Info("DeleteMediaProduct finished successfully")
	c.JSON(200, gin.H{"message": "Photo deleted successfully"})
}

// @Summary UpdateLimitOfProduct
// @Security ApiKeyAuth
// @Description Update Limit Of Product
// @Tags PRODUCTS
// @Param product_id path string true "product_id"
// @Param info body models.LimitOfProductRequest true "info"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /products/limit/{product_id} [put]
func (h *newProducts) UpdateLimitOfProduct(c *gin.Context) {
	h.Log.Info("UpdateLimitOfProduct called")
	id := c.Param("product_id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		h.Log.Error("Product ID is required")
		return
	}
	var req models.LimitOfProductRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Log.Error(err.Error())
		return
	}
	_, err = h.Product.UpdateLimitOfProduct(c, &pb.UpdateLimitOfProductRequest{Id: id, LimitOfProduct: req.LimitOfProductRequestType})
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating product limit"})
		return
	}
	h.Log.Info("UpdateLimitOfProduct finished successfully")
	c.JSON(200, gin.H{"message": "Product limit updated successfully"})
}
