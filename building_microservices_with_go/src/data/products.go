package data

import (
	"encoding/json"
	"io"
	"time"
)

// ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

type Product struct {
	// using struct tags with->json:"new_name"
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// New Constructor
func New(
	id int, name string, description string,
	price float32, sku string, createdOn string,
	updatedOn string, deletedOn string) *Product {
	
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		SKU:         sku,
		CreatedOn:   createdOn,
		UpdatedOn:   updatedOn,
		DeletedOn:   deletedOn,
	}
}

func GetProducts() ProductList {
	return productList
}

// ProductList defining a list of Product type
type ProductList []*Product

func (receiver *ProductList) ToJson(ioWriter io.Writer) error {
	encoder := json.NewEncoder(ioWriter)
	// encoding self in the function
	return encoder.Encode(receiver)
}

// ☰☰☰☰☰☰☰☰☰☰☰☰☰ DUMMY-DATA ☰☰☰☰☰☰☰☰☰☰☰☰☰
var productList = ProductList{
	New(
		1,
		"Latte",
		"Frothy milky coffee",
		2.45,
		"abc123",
		time.Now().UTC().String(),
		time.Now().UTC().String(),
		"",
	),
	New(
		2,
		"Espresso",
		"Short & strong coffee with milk",
		1.99,
		"fjd34",
		time.Now().UTC().String(),
		time.Now().UTC().String(),
		"",
	),
}

// ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
