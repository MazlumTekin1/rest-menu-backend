package gorm

//PATH: internal/adapters/output/gorm/category_repository.go

import (
	"context"
	"rest-menu-service/internal/domain"
	output "rest-menu-service/internal/ports/output"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) output.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryRepository) GetByID(ctx context.Context, id int) (*domain.Category, error) {
	var category domain.Category
	err := r.db.WithContext(ctx).
		Preload("Categories.Products").
		Where("category_id = ? AND is_deleted = ?", id, false).
		First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).
		Where("category_id = ? AND is_deleted = ?", category.ID, false).
		Updates(category).Error
}

func (r *categoryRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).
		Model(&domain.Category{}).
		Where("category_id = ? AND is_deleted = ?", id, false).
		Updates(map[string]interface{}{
			"is_deleted": true,
		}).Error
}

func (r *categoryRepository) List(ctx context.Context, filter output.CategoryFilter) ([]domain.Category, error) {
	var categorys []domain.Category
	query := r.db.WithContext(ctx).
		Preload("Categories", "is_deleted = ?", false).
		Preload("Categories.Products", "is_deleted = ?", false).
		Where("is_deleted = ?", false)

	if filter.MenuID > 0 {
		query = query.Where("restaurant_id = ?", filter.MenuID)
	}

	err := query.Find(&categorys).Error
	return categorys, err
}
