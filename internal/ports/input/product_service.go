package ports

import (
	"context"
	"rest-menu-service/internal/application/commands"
	"rest-menu-service/internal/application/dto"
	"rest-menu-service/internal/application/queries"
)

type ProductService interface {
	CreateProduct(ctx context.Context, cmd commands.CreateProductCommand) (*dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, cmd commands.UpdateProductCommand) error
	DeleteProduct(ctx context.Context, cmd commands.DeleteProductCommand) error
	GetProduct(ctx context.Context, query queries.GetProductQuery) (*dto.ProductResponse, error)
	ListProducts(ctx context.Context, query queries.ListProductsQuery) ([]dto.ProductResponse, error)
}
