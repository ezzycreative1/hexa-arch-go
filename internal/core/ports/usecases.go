package ports

import (
	"github.com/ezzycreative1/hexa-arch/internal/core/domain"
)

type TodoUseCase interface {
	Get(id string) (*domain.ToDo, error)
	List() ([]domain.ToDo, error)
	Create(title, description string) (*domain.ToDo, error)
}
