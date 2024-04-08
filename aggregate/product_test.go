package aggregate_test

import (
	"testing"

	"github.com/fabriciomoeller/ddd-go/aggregate"
)

func TestProduct_NewProduct(t *testing.T) {

	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "valid product",
			name:        "teste de novo Produto",
			description: "Descrição do produto",
			price:       10.0,
			expectedErr: nil,
		},
		{
			test:        "produto invalido",
			name:        "",
			expectedErr: aggregate.ErrMissingValues,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewProduct(tc.name, tc.description, tc.price)
			if err != tc.expectedErr {
				t.Errorf("Erro esperado %v, obtido %v", err, tc.expectedErr)
			}
		})
	}
}
