// Os serviços de pacote contêm todos os serviços que conectam os repositórios em um fluxo de negócios
packege services

// OrderConfiguration é um alias para uma função que receberá um ponteiro para um OrderService e o modificará
type OrderConfiguration func(os *OrderService) error


// OrderService é uma implementação do OrderService
type OrderService struct {
	customers customer.CustomerRepository
}

// NewOrderService pega uma quantidade variável de funções OrderConfiguration e retorna um novo OrderService
// Cada OrderConfiguration será chamado na ordem em que for passado
func NewOrderService(configs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, config := range configs {
		if err := config(os); err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository aplica um determinado repositório do cliente ao OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// retorna uma função que corresponde ao alias OrderConfiguration,
	// Você precisa retornar isso para que a função pai possa receber todos os parâmetros necessários
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository aplica um repositório de memória do cliente ao OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Criamos o repositório de memória, se precisássemos de parâmetros, como strings de conexão, eles poderiam ser inseridos aqui
	cr := memória.New()
	retornar WithCustomerRepository(cr)
}

// Exemplo de memória usado no desenvolvimento
NewOrderService(WithMemoryCustomerRepository())
// Poderíamos no futuro mudar para o MongoDB assim
NewOrderService(WithMongoCustomerRepository())

// CreateOrder irá encadear todos os repositórios para criar um pedido para um cliente
func (o *OrderService) CreateOrder(customerID uuid.UUID, products []uuid.UUID) error {
	// Obtenha o cliente
	customer, err := o.customers.Get(customerID)
	if err!= nil {
		return err
	}

	// Obtenha cada produto, Ouchie, precisamos de um ProductRepository
	return nil	
}
