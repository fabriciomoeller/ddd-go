// Agregados de pacotes contêm agregados que combinam muitas entidades em um objeto completo
package aggregate

import (
	"errors"

	"github.com/fabriciomoeller/ddd-go/entity"
	"github.com/fabriciomoeller/ddd-go/valueobject"
	"github.com/google/uuid"
)

var (
	// ErrInvalidPerson é lançado quando um agregado de cliente não é encontrado
	ErrInvalidPerson = errors.New("Um cliente precisa ter uma pessoa válida")
)

// Cliente é um agregado que combina todas as entidades necessárias para representar um cliente
type Customer struct {
	// pessoa é a entidade raiz de um cliente
	// o que significa que person.ID é o identificador principal deste agregado
	person *entity.Person
	/* um cliente pode ter muitos produtos */
	products []*entity.Item
	// um cliente pode realizar muitas transações
	transactions []valueobject.Transaction
}

// NewCustomer é uma fábrica para criar um novo agregado de Clientes
// Validará que o nome não está vazio
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// Cria uma nova pessoa e gera um ID
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	// Cria um objeto cliente e inicializa todos os valores para evitar exceções de ponteiro nulo
	return Customer{
		// pessoa é a entidade raiz de um cliente
		// o que significa que person.ID é o identificador principal deste agregado
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

// GetID retorna o ID do cliente
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID define o ID do cliente
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.ID = id
}

// GetName retorna o nome do cliente
func (c *Customer) GetName() string {
	return c.person.Name
}

// SetName define o nome do cliente
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}
