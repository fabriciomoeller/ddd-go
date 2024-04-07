package entity

import (
	"github.com/google/uuid"
)

// Person é uma entidade que representa uma pessoa em todos os Domínios
type Person struct {
	// ID é o identificador da Entidade, o ID é compartilhado por todos os subdomínios
	ID   uuid.UUID
	Name string
	Age  int
}
