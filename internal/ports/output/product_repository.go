package ports

//PATH: internal/ports/output/product_repository.go

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

type ProductReader interface {
	GetByID(ctx context.Context, id int) (*domain.Product, error)
	List(ctx context.Context, filter ProductFilter) ([]domain.Product, error)
	ListWithDetails(ctx context.Context, filter ProductFilter) ([]dto.ProductResponse, error)
}

type ProductWriter interface {
	Create(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id int, updateUserID int) error
}

type ProductRepository interface {
	ProductReader
	ProductWriter
}
