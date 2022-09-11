package routes

import (
	"context"
	"fmt"
	"os"

	"github.com/YogaPratama02/go-crud-mongo/config"
	"github.com/YogaPratama02/go-crud-mongo/controllers"
	"github.com/YogaPratama02/go-crud-mongo/repositories"
	"github.com/YogaPratama02/go-crud-mongo/services"
	"github.com/gin-gonic/gin"
)

func Init() {
	db := config.ConnectDB()
	ctx := context.TODO()
	// defer db.Close()

	// PRODUCT
	productRepository := repositories.NewProductRepository(db, ctx)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Recovery())
	v1 := e.Group("/api/v1")

	product := v1.Group("product")
	{
		product.POST("", productController.CreateProductController)
		product.GET("", productController.GetProductController)
		product.PUT("", productController.UpdateProductController)
		product.DELETE("", productController.DeleteProductController)

	}

	e.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
