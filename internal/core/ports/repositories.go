package ports

import (
	"github.com/ezzycreative1/hexa-arch/internal/core/domain"
)

type TodoRepository interface {
	Get(id string) (*domain.ToDo, error)
	List() ([]domain.ToDo, error)
	Create(todo *domain.ToDo) (*domain.ToDo, error)
}
