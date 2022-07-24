package intraprocess

import (
	"github.com/leksyking/monolith-microservice/pkg/orders/application"
	"github.com/leksyking/monolith-microservice/pkg/orders/domain/orders"
)

type OrdersInterface struct {
	sevice application.OrdersService
}

func NewOrdersInterface(service application.OrdersService) OrdersInterface {
	return OrdersInterface{}
}
func (p OrdersInterface) MarkAsPaid(orderID, string) error {
	return p.service.MarkOrderAsPaid(application.MarkOrderAsPaid{orders.ID(orderID)})

}
