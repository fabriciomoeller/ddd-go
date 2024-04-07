package customer

import (
	"errors"

	"github.com/fabriciomoeller/ddd-go/aggregate"
	"github.com/google/uuid"
)

var (
	// ErrCustomerNotFound é retornado quando um cliente não é encontrado.
	ErrCustomerNotFound = errors.New("o cliente não foi encontrado no repositório")
	// ErrFailedToAddCustomer é retornado quando o cliente não pôde ser adicionado ao repositório.
	ErrFailedToAddCustomer = errors.New("falha ao adicionar o cliente ao repositório")
	// ErrUpdateCustomer é retornado quando o cliente não pôde ser atualizado no repositório.
	ErrUpdateCustomer = errors.New("falha ao atualizar o cliente no repositório")
)

// CustomerRepository é uma interface que define as regras sobre o que um repositório do cliente
// Tem que ser capaz de executar
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
	Delete(uuid.UUID) error
	List() ([]aggregate.Customer, error)
}
