package ports

//PATH: internal/ports/input/category_service.go

import (
	"context"
	"rest-menu-service/internal/application/commands"
	"rest-menu-service/internal/application/dto"
	"rest-menu-service/internal/application/queries"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, cmd commands.CreateCategoryCommand) (*dto.CategoryResponse, error)
	UpdateCategory(ctx context.Context, cmd commands.UpdateCategoryCommand) error
	DeleteCategory(ctx context.Context, cmd commands.DeleteCategoryCommand) error
	GetCategory(ctx context.Context, query queries.GetCategoryQuery) (*dto.CategoryResponse, error)
	ListCategorys(ctx context.Context, query queries.ListCategoriesQuery) ([]dto.CategoryResponse, error)
}
