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
		card.GET("/amount/:card_id", h.CardsMethods().GetAmountOfUserCard)
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
		prcs.POST("")
		prcs.GET("/products/:product_id")
		prcs.GET("/user/:product_id/:user_id")
		prcs.GET("")
		prcs.GET("/:id")
		prcs.PUT("/:id")
		prcs.DELETE("/:id")

	}

	wish := router.Group("/wishlist")
	{
		wish.POST("")
		wish.GET("")
		wish.GET("/:id")
	}

	fdbk := router.Group("/feedback")
	{
		fdbk.POST("/:product_id")
		fdbk.GET("")
		fdbk.GET("/:product_id")

	}

	bght := router.Group("/bought")
	{
		bght.GET("/:product_id")
		bght.GET("")
	}

	ntfc := router.Group("/notifications")
	{
		ntfc.GET("/all")
		ntfc.GET("/unreaden")
	}

	return router
}
