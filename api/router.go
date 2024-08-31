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
func Router(hand *handler.Handler) *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.Check)
	router.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))

	card := router.Group("/cards")
	{
		card.POST("",)
		card.GET("",)
		card.GET("/amount",)
	}

	prd := router.Group("/products")
	{
		prd.POST("", )
		prd.POST("/list", )
		prd.GET("/:id", )
		prd.PUT("/:id", )
		prd.DELETE("/:id", )
		prd.POST("/photo/:product_id", )
		prd.DELETE("/photo/:product_id", )

	}

	prcs := router.Group("/process")
	{
		prcs.POST("", )
		prcs.GET("/:prduct_id",)
		prcs.GET("/:prduct_id/:user_id",)
		prcs.GET("",)
		prcs.GET("/:id",)
		prcs.PUT("/:id",)
		prcs.DELETE("/:id")

	}

	wish := router.Group("/wishlist")
	{
		wish.POST("", )
		wish.GET("", )
		wish.GET("/:id", )
	}

	fdbk := router.Group("/feedback")
	{
		fdbk.POST("/:product_id", )
		fdbk.GET("", )
		fdbk.GET("/:product_id", )

	}

	bght := router.Group("/bought")
	{
		bght.GET("/:product_id", )
		bght.GET("", )
	}


	return router
}
