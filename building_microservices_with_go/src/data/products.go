package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

// Product defines the structure for an API product
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

func NewProduct(id int,
	name string,
	description string,
	price float32,
	sku string,
	createdOn string,
	updatedOn string,
	deletedOn string) *Product {

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

func (receiver *ProductList) ToJson(ioWriter io.Writer) error {
	encoder := json.NewEncoder(ioWriter)
	// encoding self in the function
	return encoder.Encode(receiver)
}

func (receiver *Product) FromJson(jsonIOReader io.Reader) error {
	// NewDecoder returns a new decoder that reads from the reader
	encoder := json.NewDecoder(jsonIOReader)
	return encoder.Decode(receiver)
}

// ProductList defining a list of Product type
type ProductList []*Product

/* HTTP-VERB functions */

func GetProducts() ProductList {
	return productList
}

func AddProducts(product *Product) {
	product.ID = getNextID()
	// add the ID to our product list
	productList = append(productList, product)
}

func UpdateProducts(id int, product *Product) error {
	_, position, err := findProduct(id)
	if err != nil {
		return err
	}

	product.ID = id
	productList[position] = product

	return nil
}

// ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ utility functions ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
func getNextID() int {
	tailEnd := len(productList) - 1
	lastProductInTheList := productList[tailEnd]
	return lastProductInTheList.ID + 1
}

func findProduct(id int) (*Product, int, error) {
	for i, product := range productList {
		if product.ID == id {
			return product, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

var ErrProductNotFound = fmt.Errorf("product not found")

// ☰☰☰☰☰☰☰☰☰☰☰☰☰ DUMMY-DATA ☰☰☰☰☰☰☰☰☰☰☰☰☰
var productList = ProductList{
	NewProduct(
		1,
		"Latte",
		"Frothy milky coffee",
		2.45,
		"abc123",
		time.Now().UTC().String(),
		time.Now().UTC().String(),
		"",
	),
	NewProduct(
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
