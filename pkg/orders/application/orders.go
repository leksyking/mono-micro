package application

type productsService interface {
}

type paymentsService interface {
}

type ordersService struct {
}

func NewOrdersService() {

}

type PlaceOrderCommand struct {
}

func (s ordersService) PlaceOrder(cmd PlaceOrderCommand) error {

}

type MarkOrderAsPaidCommand struct {
}

func (s ordersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {

}

func (s ordersService) OrderByID(id orders.ID) (orders.Order, error) {

}
