package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/leksyking/monolith-microservice/pkg/common/price"
)

func AddRoutes(router *chi.Mux, productsReadModel ProductReadModel) {
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

type ProductReadModel interface {
	AllProducts([]products.Product, error)
}

type productsResource struct {
	readModel ProductReadModel
}

type productView struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       priceView `json:"price"`
}

type priceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func priceViewFromPrice(p price.Price) priceView {
	return priceView{p.Cents(), p.Currency()}
}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.readModel.AllProducts()
	if err != nil {
		render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	view := []productView{}
	for _, product := range products {
		view = append(view, productView{
			string(product.ID()),
			product.Name(),
			product.Description(),
			priceViewFromPrice(product.Price()),
		})
	}
	render.Respond(w, r, view)
}
