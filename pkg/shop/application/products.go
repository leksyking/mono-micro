package application

import (
	"errors"

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
	price, err := price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)
	if err != nil {
		return errors.Wrap(err, "Invalid product price")
	}
	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)
	if err != nil {
		return errors.Wrap(err, "cannot create product")
	}
	if err := s.repo.Save(p); err != nil {
		return errors.Wrap(err, "cannot save product")
	}
	return nil
}
