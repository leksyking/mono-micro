package application

import (
	"github.com/leksyking/monolith-microservice/pkg/common/price"
	"github.com/leksyking/monolith-microservice/pkg/shop/domain/products"
)

type productReadModel interface {
	AllProducts() ([]products.Product, error)
}

type ProductsService struct {
	repo      products.Repository
	readModel productReadModel
}

func NewProductsService() ProductsService {

}

func (s ProductsService) AllProducts() ([]products.Product, error) {

}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCents    uint
	PriceCurrency string
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {
	price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)
	products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)
	s.repo.Save
}
