// A memória do pacote é uma implementação na memória do repositório do cliente
package memory

import (
	"fmt"
	"sync"

	"github.com/fabriciomoeller/ddd-go/aggregate"
	"github.com/fabriciomoeller/ddd-go/domain/customer"
	"github.com/google/uuid"
)

// MemoryRepository cumpre a interface CustomerRepository
type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// New é uma função de fábrica para gerar um novo repositório de clientes
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get encontra um cliente por ID
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	// Se o cliente não for encontrado, retorne um erro
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

// Adicionar adicionará um novo cliente ao repositório
func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		// A verificação de segurança dos clientes não é criada, não deveria acontecer se estiver usando a Fábrica, mas nunca se sabe
		// Se o repositório estiver vazio, inicialize-o com um novo mapa de clientes
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}

	// Certifique-se de que o Cliente ainda não esteja no repositório
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("cliente já existe no repositório: %w", customer.ErrFailedToAddCustomer)
	}

	// Adicione o cliente ao repositório
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil

}

// Atualizar um cliente existente no repositório
func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	mr.Lock()
	defer mr.Unlock()

	// Certifique-se de que o cliente esteja no repositório
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("cliente não existe no repositório: %w", customer.ErrUpdateCustomer)
	}

	// Atualize o cliente no repositório
	mr.customers[c.GetID()] = c

	return nil
}

// Remover removerá um cliente do repositório
func (mr *MemoryRepository) Remove(id uuid.UUID) error {
	return nil
}

// List retornará uma lista de todos os clientes no repositório
func (mr *MemoryRepository) List() ([]aggregate.Customer, error) {
	return []aggregate.Customer{}, nil
}
