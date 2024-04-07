package memory

import (
	"testing"

	"github.com/fabriciomoeller/ddd-go/aggregate"
	"github.com/fabriciomoeller/ddd-go/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	// Cria um cliente falso para adicionar ao repositório
	cust, err := aggregate.NewCustomer("Pedro Quinotto Moeller")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()
	// Cria o repo para uso, e adiciona dados para teste
	// Ignora a fábrica para isso
	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}
	testCases := []testCase{
		{
			name:        "Cliente encontrado com sucesso",
			id:          id,
			expectedErr: nil,
		},
		{
			name:        "Client não encontrado",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Erro esperado %v, obtido %v", err, tc.expectedErr)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Cliente adicionado com sucesso",
			cust:        "Pedro Quinotto Moeller",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]aggregate.Customer{},
			}
			cust, err := aggregate.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}
			err = repo.Add(cust)
			if err != tc.expectedErr {
				t.Errorf("Erro esperado %v, obtido %v", err, tc.expectedErr)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetName() != cust.GetName() {
				t.Errorf("Nome esperado %s, obtido %s", tc.cust, found.GetName())
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("ID esperado %s, obtido %s", cust.GetID(), found.GetID())
			}
		})
	}
}
