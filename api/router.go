package api

import (
	_ "api/api/docs"
	"api/api/handler"
	"api/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @title All
// @version 1.0
// @description API Gateway
// BasePath: /
func Router(h handler.HandlerInterface) *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.Check)
	router.Use(h.EnforcerMethods().CheckPermissionMiddleware())

	card := router.Group("/cards")
	{
		card.POST("", h.CardsMethods().CreateCards)
		card.GET("", h.CardsMethods().GetUserCards)
		card.GET("/amount/:card_number", h.CardsMethods().GetAmountOfUserCard)
	}

	prd := router.Group("/products")
	{
		prd.POST("", h.ProductMethods().CreateProduct)
		prd.POST("/list", h.ProductMethods().GetProductsList)
		prd.GET("/:id", h.ProductMethods().GetProductByID)
		prd.GET("", h.ProductMethods().GetSellerProducts)
		prd.PUT("/:id", h.ProductMethods().UpdateProduct)
		prd.DELETE("/:id", h.ProductMethods().DeleteProduct)
		prd.POST("/photo/:product_id", h.ProductMethods().UploadProductPhoto)
		prd.DELETE("/photo/:product_id", h.ProductMethods().DeleteProductPhoto)
	}

	prcs := router.Group("/process")
	{
		prcs.POST("/buy", h.ProcessMethods().CreateProcess)
		prcs.GET("/products/:product_id", h.ProcessMethods().GetProcessByProductId)
		prcs.GET("/user/:product_id/:user_id", h.ProcessMethods().GetProcessOfUserByProductId)
		prcs.GET("", h.ProcessMethods().GetProcessByUserId)
		prcs.GET("/:id", h.ProcessMethods().GetProcessById)
		prcs.PUT("/:id", h.ProcessMethods().UpdateProcess)
		prcs.DELETE("/:id", h.ProcessMethods().CancelProcess)

	}

	wish := router.Group("/wishlist")
	{
		wish.POST("/:product_id", h.WishlistMethods().CreateWishlist)
		wish.GET("", h.WishlistMethods().GetWishlist)
		wish.GET("/:id", h.WishlistMethods().GetWishlistById)
	}

	fdbk := router.Group("/feedback")
	{
		fdbk.POST("/:product_id", h.FeedbackMethods().CreateFeedback)
		fdbk.GET("", h.FeedbackMethods().GetFeedbackOfUser)
		fdbk.GET("/:product_id", h.FeedbackMethods().GetFeedback)
	}

	bght := router.Group("/bought")
	{
		bght.GET("/:product_id", h.BoughtMethods().GetBought)
		bght.GET("", h.BoughtMethods().GetBoughtOfUser)
	}

	ntfc := router.Group("/notifications")
	{
		ntfc.GET("/all")
		ntfc.GET("/unreaden")
	}

	return router
}
