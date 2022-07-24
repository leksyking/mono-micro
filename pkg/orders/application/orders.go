package application

import (
	"log"

	"github.com/leksyking/monolith-microservice/pkg/common/price"
	"github.com/leksyking/monolith-microservice/pkg/orders/domain/orders"
	"github.com/pkg/errors"
)

type productsService interface {
	ProductsByID(id orders.ProductID) (orders.Product, error)
}

type paymentsService interface {
	InitializeOrderPayment(id orders.ID, price price.Price) error
}

type OrdersService struct {
	productsService  productsService
	paymentsService  paymentsService
	ordersRepository orders.Repository
}

func NewOrdersService(productsService productsService, paypaymentsService paymentsService, ordersRepository ordersRepository) OrdersService {
	return OrdersService{productsService, paypaymentsService, ordersRepository}
}

type PlaceOrderCommand struct {
	OrderID   orders.ID
	ProductID orders.ProductID
}

type PlaceOrderCommandAddress struct {
	Name       string
	Street     string
	City       string
	PostalCode string
	Country    string
}

func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {
	address, err := orders.NewAddress(
		cmd.Address.Name,
		cmd.Address.Street,
		cmd.Address.City,
		cmd.Address.PostalCode,
		cmd.Address.Country,
	)
	if err != nil {
		return errors.Wrap(err, "invalid address")
	}

	product, err := s.productsService.ProductByID(cmd.ProductID)
	if err != nil {
		return errors.Wrapf(err, "cannot get the product")
	}
	newOrder, err := orders.NewOrder(cmd.OrderID, product, address)
	if err != nil {
		return errors.Wrap("Cannot create order")
	}
	if err := s.ordersRepository.Save(newOrder); err != nil {
		return errors.Wrap(err, "cannot save order")
	
	if err := s.paymentsService.InitializeOrderPayment(newOrder.ID(), newOrder.Product().Price()); err != nil {
		return errors.Wrap(err, "cannot initialze order payment")
	}
	log.Printf("order %s placed", cmd.OrderID)
	return nil
}

type MarkOrderAsPaidCommand struct {
	OrdersID orders.ID
}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {
	o, err := s.orderRepository.ByID(cmd.OrderID)
	if err != nil{
		return errors.Wrapf(err, "cannot get order %s", cmd.OrderID)
	}
	o.MarkAsPaid()
	if err:= s.ordersRepository.Save(o); err !=nil{
		return errors.Wrap(err, "cannot  save order")
	}
	log.Println("marked order %s as paid", cmd.OrderID)
	return nil
}

func (s OrdersService) OrderByID(id orders.ID) (orders.Order, error) {
	o, err := s.ordersRepository.ByID(id)
	if err != nil {
		return orders.Order{}, errors.Wrapf(err, "Cannot get order %s", id)
	}
	return *o, nil
}
