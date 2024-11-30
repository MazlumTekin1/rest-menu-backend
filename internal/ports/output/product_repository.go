package ports

import (
	"context"
	"rest-menu-service/internal/application/dto"
	"rest-menu-service/internal/domain"
)

type Pagination struct {
	Page     int
	PageSize int
}

type ProductFilter struct {
	MenuID     int
	CategoryID int
	Pagination Pagination
}

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	GetByID(ctx context.Context, id int) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id int, updateUserID int) error
	List(ctx context.Context, filter ProductFilter) ([]domain.Product, error)
	ListWithDetails(ctx context.Context, filter ProductFilter) ([]dto.ProductResponse, error)
}
