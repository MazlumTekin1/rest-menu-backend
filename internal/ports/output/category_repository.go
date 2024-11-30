package ports

import (
	"context"
	"rest-menu-service/internal/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, menu *domain.Category) error
	GetByID(ctx context.Context, id int) (*domain.Category, error)
	Update(ctx context.Context, menu *domain.Category) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter CategoryFilter) ([]domain.Category, error)
}

type CategoryFilter struct {
	MenuID    int
	IsDeleted bool
}
