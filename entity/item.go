package entity

import "github.com/google/uuid"

// Item representa um Item para todos os subdomínios
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
