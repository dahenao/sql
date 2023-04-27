package product

import (
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type Repository interface {
	// GetByID busca un producto por su id
	GetByID(id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Create(p domain.Product) (domain.Product, error)
	// Update actualiza un producto
	Update(id int, p domain.Product) (domain.Product, error)
	// Delete elimina un producto
	Delete(id int) error
	// mostrar todos los registros de la tabla
}
