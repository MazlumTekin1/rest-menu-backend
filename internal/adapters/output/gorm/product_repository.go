package gorm

//PATH: internal/adapters/output/gorm/product_repository.go

import (
	"context"
	"rest-menu-service/internal/application/dto"
	"rest-menu-service/internal/domain"
	output "rest-menu-service/internal/ports/output"
	"time"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) output.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepository) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	var product domain.Product

	err := r.db.WithContext(ctx).
		Select("product_id, menu_id, category_id, product_name, product_price, product_description, product_image_url, created_date, updated_date, is_deleted, create_user_id, update_user_id").
		Where("product_id = ? AND is_deleted = ?", id, false).
		First(&product).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, output.ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).
		Where("product_id = ? AND is_deleted = ?", product.ID, false).
		Updates(product).Error
}

func (r *productRepository) Delete(ctx context.Context, id int, updateUserID int) error {
	result := r.db.WithContext(ctx).
		Model(&domain.Product{}).
		Where("product_id = ? AND is_deleted = ?", id, false).
		Updates(map[string]interface{}{
			"is_deleted":     true,
			"updated_date":   time.Now(),
			"update_user_id": updateUserID,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return output.ErrProductNotFound
	}

	return nil
}

func (r *productRepository) List(ctx context.Context, filter output.ProductFilter) ([]domain.Product, error) {
	var products []domain.Product

	query := r.db.WithContext(ctx).
		Select("products.product_id, products.menu_id, products.category_id, products.product_name, "+
			"products.product_price, products.product_description, products.product_image_url, "+
			"products.created_date, products.updated_date, products.is_deleted, "+
			"products.create_user_id, products.update_user_id").
		Where("products.is_deleted = ?", false)

	// MenuID filtresi
	if filter.MenuID > 0 {
		query = query.Where("products.menu_id = ?", filter.MenuID)
	}

	// CategoryID filtresi
	if filter.CategoryID > 0 {
		query = query.Where("products.category_id = ?", filter.CategoryID)
	}

	err := query.Order("products.product_id asc").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) ListWithDetails(ctx context.Context, filter output.ProductFilter) ([]dto.ProductResponse, error) {
	var products []dto.ProductResponse

	// Base query
	query := r.db.WithContext(ctx).
		Table("rest_user.products as p").
		Select(`
            p.product_id as id,
            p.product_name as name,
            p.product_price as price,
            p.product_description as description,
            p.product_image_url as image_url,
            p.menu_id,
            rm.menu_name,
            p.category_id,
            c.category_name,
            p.created_date,
            p.updated_date
        `).
		Joins("LEFT JOIN rest_user.restaurant_menus rm ON rm.menu_id = p.menu_id").
		Joins("LEFT JOIN rest_user.categories c ON c.category_id = p.category_id").
		Where("p.is_deleted = ?", false)

	// Apply filters
	if filter.MenuID > 0 {
		query = query.Where("p.menu_id = ?", filter.MenuID)
	}

	if filter.CategoryID > 0 {
		query = query.Where("p.category_id = ?", filter.CategoryID)
	}

	// Apply pagination
	if filter.Pagination.Page > 0 && filter.Pagination.PageSize > 0 {
		offset := (filter.Pagination.Page - 1) * filter.Pagination.PageSize
		query = query.Offset(offset).Limit(filter.Pagination.PageSize)
	}

	// Execute query
	err := query.Order("p.product_id asc").Scan(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
