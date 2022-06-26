package products

type MemoryRepository struct {
	product []products.Product
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{[]products.Product{}}
}
func (m *MemoryRepository) Save(productToSave *products.Product) error {
	for i, p := range m.product {
		if p.ID() == productToSave.ID() {
			m.product[i] = *productToSave
			return nil
		}
	}
	m.product = append(m.product, *productToSave)
	return nil
}

func (m *MemoryRepository) ByID(id products.ID) (*products.Product, error) {
	for _, p := range m.product {
		if p.ID() == id {
			return &p, nil
		}
	}
	return nil, products.ErrNotFound
}

func (m *MemoryRepository) AllProducts() ([]products.Product, error) {
	return m.product, nil
}
