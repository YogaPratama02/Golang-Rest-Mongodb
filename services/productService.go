package services

import (
	"log"
	"time"

	"github.com/YogaPratama02/go-crud-mongo/models"
	"github.com/YogaPratama02/go-crud-mongo/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProduct() (*[]models.Product, error)
	UpdateProduct(Id string, product *models.Product) error
	DeleteProduct(Id string) error
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) *productService {
	return &productService{repository}
}

func (s *productService) CreateProduct(product *models.Product) error {
	productCreate := models.Product{
		ProductName:  product.ProductName,
		SKU:          product.SKU,
		PriceProduct: product.PriceProduct,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := s.productRepository.CreateProduct(&productCreate); err != nil {
		log.Printf("Error create product with err: %s", err)
		return err
	}

	return nil
}

func (s *productService) GetProduct() (*[]models.Product, error) {
	products, err := s.productRepository.GetProduct()
	if err != nil {
		log.Printf("Error get data all products with err: %s", err)
		return nil, err
	}

	return products, nil
}

func (s *productService) UpdateProduct(Id string, product *models.Product) error {
	obId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		log.Printf("Error convert id to object id err: %s", err)
		return err
	}

	productUpdate := models.Product{
		Id:           obId,
		ProductName:  product.ProductName,
		SKU:          product.SKU,
		PriceProduct: product.PriceProduct,
		UpdatedAt:    time.Now(),
	}
	if err := s.productRepository.UpdateProduct(&productUpdate); err != nil {
		log.Printf("Error update product with err: %s", err)
		return err
	}
	return nil
}

func (s *productService) DeleteProduct(Id string) error {
	obId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		log.Printf("Error convert id to object id err: %s", err)
		return err
	}

	err = s.productRepository.DeleteProduct(obId)
	if err != nil {
		log.Printf("Error delete product with err: %s", err)
		return err
	}

	return nil
}
