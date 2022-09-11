package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductName  string             `json:"product_name,omitempty" bson:"product_name,omitempty" validate:"required"`
	SKU          string             `json:"sku,omitempty" bson:"sku,omitempty" validate:"required"`
	PriceProduct float64            `json:"price_product,omitempty" bson:"price_product,omitempty" validate:"required"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
