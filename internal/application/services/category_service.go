package services

import (
	"context"
	"rest-menu-service/internal/application/commands"
	"rest-menu-service/internal/application/dto"
	"rest-menu-service/internal/application/queries"
	"rest-menu-service/internal/domain"
	output "rest-menu-service/internal/ports/output"
	"time"
)

type categoryService struct {
	categoryRepo output.CategoryRepository
}

func NewCategoryService(categoryRepo output.CategoryRepository) *categoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, cmd commands.CreateCategoryCommand) (*dto.CategoryResponse, error) {
	category, err := domain.NewCategory(cmd.Name, cmd.MenuID, cmd.CreateUserID)
	if err != nil {
		return nil, err
	}

	if err := s.categoryRepo.Create(ctx, category); err != nil {
		return nil, err
	}

	return mapToCategoryResponse(category), nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, cmd commands.UpdateCategoryCommand) error {
	category, err := s.categoryRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	category.Name = cmd.Name
	category.UpdatedDate = time.Now()

	return s.categoryRepo.Update(ctx, category)
}

func (s *categoryService) DeleteCategory(ctx context.Context, cmd commands.DeleteCategoryCommand) error {
	if _, err := s.categoryRepo.GetByID(ctx, cmd.ID); err != nil {
		return err
	}

	return s.categoryRepo.Delete(ctx, cmd.ID)
}

func (s *categoryService) GetCategory(ctx context.Context, query queries.GetCategoryQuery) (*dto.CategoryResponse, error) {
	category, err := s.categoryRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	return mapToCategoryResponse(category), nil
}

func (s *categoryService) ListCategorys(ctx context.Context, query queries.ListCategoriesQuery) ([]dto.CategoryResponse, error) {
	filter := output.CategoryFilter{
		MenuID: query.MenuID,
	}

	categorys, err := s.categoryRepo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	response := make([]dto.CategoryResponse, len(categorys))
	for i, category := range categorys {
		response[i] = *mapToCategoryResponse(&category)
	}

	return response, nil
}

func mapToCategoryResponse(category *domain.Category) *dto.CategoryResponse {
	return &dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		MenuID:    category.MenuID,
		CreatedAt: category.CreatedDate.Format(time.RFC3339),
		UpdatedAt: category.UpdatedDate.Format(time.RFC3339),
	}
}
