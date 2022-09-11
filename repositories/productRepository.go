package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/YogaPratama02/go-crud-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProduct() (*[]models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(Id primitive.ObjectID) error
}

type productRepository struct {
	db  *mongo.Database
	ctx context.Context
}

func NewProductRepository(db *mongo.Database, ctx context.Context) *productRepository {
	return &productRepository{db, ctx}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	db := r.db
	_, err := db.Collection("products").InsertOne(r.ctx, &product)
	if err != nil {
		log.Printf("Error create product to database with err: %s", err)
		return err
	}
	// res.InsertedID()
	return nil
}

func (r *productRepository) GetProduct() (*[]models.Product, error) {
	db := r.db
	var products = []models.Product{}
	cur, err := db.Collection("products").Find(r.ctx, bson.M{})
	if err != nil {
		log.Printf("Error get all products to database with err: %s", err)
		return nil, err
	}
	for cur.Next(r.ctx) {
		product := models.Product{}
		err := cur.Decode(&product)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return &products, nil
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	db := r.db
	query := bson.M{"_id": product.Id}
	update := bson.D{{Key: "$set", Value: product}}
	err := db.Collection("products").FindOneAndUpdate(r.ctx, query, update).Decode(&product)

	if err != nil {
		log.Printf("Error update product to database with err: %s", err)
		return err
	}
	return nil
}

func (r *productRepository) DeleteProduct(Id primitive.ObjectID) error {
	db := r.db
	query := bson.M{"_id": Id}
	res, err := db.Collection("products").DeleteOne(r.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("product Id not fount")
	}
	return nil
}
