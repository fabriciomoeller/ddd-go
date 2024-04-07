package aggregate_test

import (
	"errors"
	"testing"

	"github.com/fabriciomoeller/ddd-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {

	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	// Cria novos casos de teste
	testCases := []testCase{
		{
			test:        "Nome invalido por estar vazio",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Nome valido",
			name:        "Fabricio Moeller",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			// Cria um novo cliente

			_, err := aggregate.NewCustomer(tc.name)
			//Verifica se o erro corresponde ao erro esperado
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Novo cliente retornou %q, %q", err, tc.expectedErr)

			}
		})
	}
}
