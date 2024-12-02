package gorm

//PATH: internal/adapters/output/gorm/menu_repository.go

import (
	"context"
	"rest-menu-service/internal/domain"
	output "rest-menu-service/internal/ports/output"

	"gorm.io/gorm"
)

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) output.MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) Create(ctx context.Context, menu *domain.Menu) error {
	return r.db.WithContext(ctx).Create(menu).Error
}

func (r *menuRepository) GetByID(ctx context.Context, id int) (*domain.Menu, error) {
	var menu domain.Menu
	err := r.db.WithContext(ctx).
		Preload("Categories.Products").
		Where("menu_id = ? AND is_deleted = ?", id, false).
		First(&menu).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) Update(ctx context.Context, menu *domain.Menu) error {
	return r.db.WithContext(ctx).
		Where("menu_id = ? AND is_deleted = ?", menu.ID, false).
		Updates(menu).Error
}

func (r *menuRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).
		Model(&domain.Menu{}).
		Where("menu_id = ? AND is_deleted = ?", id, false).
		Updates(map[string]interface{}{
			"is_deleted": true,
		}).Error
}

func (r *menuRepository) List(ctx context.Context, filter output.MenuFilter) ([]domain.Menu, error) {
	var menus []domain.Menu
	query := r.db.WithContext(ctx).
		Preload("Categories", "is_deleted = ?", false).
		Preload("Categories.Products", "is_deleted = ?", false).
		Where("is_deleted = ?", false)

	if filter.RestaurantID > 0 {
		query = query.Where("restaurant_id = ?", filter.RestaurantID)
	}

	err := query.Find(&menus).Error
	return menus, err
}
