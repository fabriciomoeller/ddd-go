package aggregate

import (
	"errors"

	"github.com/fabriciomoeller/ddd-go/entity"
	"github.com/google/uuid"
)

var (
	// ErrMissingValues é retornado quando um produto é criado sem nome ou descrição
	ErrMissingValues = errors.New("valores faltantes")
)

// Produto é um agregado que combina item com preço e quantidade
type Product struct {
	// item é a entidade raiz que é um item
	item  *entity.Item
	price float64
	// Quantidade é a quantidade de produtos em estoque
	quantity int
}

// NewProduct criará um novo produto
// retornará erro se o nome da descrição estiver vazio
func NewProduct(name, descripton string, price float64) (Product, error) {
	if name == "" || descripton == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: descripton,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
