package controllers

import (
	"github.com/YogaPratama02/go-crud-mongo/helpers"
	"github.com/YogaPratama02/go-crud-mongo/models"
	"github.com/YogaPratama02/go-crud-mongo/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service}
}

func (productController *ProductController) CreateProductController(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := helpers.DoValidation(&product); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	if err := productController.service.CreateProduct(&product); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	helpers.NewHandlerResponse("Successfully create product", nil).SuccessCreate(c)
}

func (productController *ProductController) GetProductController(c *gin.Context) {
	products, err := productController.service.GetProduct()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	helpers.NewHandlerResponse("Successfully get all products", products).Success(c)
}

func (productController *ProductController) UpdateProductController(c *gin.Context) {
	var product models.Product

	id := c.Query("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := helpers.DoValidation(&product); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	if err := productController.service.UpdateProduct(id, &product); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update product", nil).Success(c)
}

func (productController *ProductController) DeleteProductController(c *gin.Context) {
	id := c.Query("id")

	if err := productController.service.DeleteProduct(id); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	helpers.NewHandlerResponse("Successfully delete product", nil).Success(c)
}
