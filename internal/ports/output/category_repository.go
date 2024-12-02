package ports

//PATH: internal/ports/output/category_repository.go

import (
	"context"
	"rest-menu-service/internal/domain"
)

type CategoryRepository interface {
	CategoryReader
	CategoryWriter
}

type CategoryFilter struct {
	MenuID    int
	IsDeleted bool
}

type CategoryReader interface {
	GetByID(ctx context.Context, id int) (*domain.Category, error)
	List(ctx context.Context, filter CategoryFilter) ([]domain.Category, error)
}

type CategoryWriter interface {
	Create(ctx context.Context, menu *domain.Category) error
	Update(ctx context.Context, menu *domain.Category) error
	Delete(ctx context.Context, id int) error
}
