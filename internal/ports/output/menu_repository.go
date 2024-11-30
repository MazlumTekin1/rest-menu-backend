package ports

import (
	"context"
	"rest-menu-service/internal/domain"
)

type MenuRepository interface {
	Create(ctx context.Context, menu *domain.Menu) error
	GetByID(ctx context.Context, id int) (*domain.Menu, error)
	Update(ctx context.Context, menu *domain.Menu) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter MenuFilter) ([]domain.Menu, error)
}

type MenuFilter struct {
	RestaurantID int
	IsDeleted    bool
}
