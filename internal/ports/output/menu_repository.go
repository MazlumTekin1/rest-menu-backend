package ports

//PATH: internal/ports/output/menu_repository.go

import (
	"context"
	"rest-menu-service/internal/domain"
)

type MenuRepository interface {
	MenuReader
	MenuWriter
}

type MenuFilter struct {
	RestaurantID int
	IsDeleted    bool
}

type MenuReader interface {
	GetByID(ctx context.Context, id int) (*domain.Menu, error)
	List(ctx context.Context, filter MenuFilter) ([]domain.Menu, error)
}

type MenuWriter interface {
	Create(ctx context.Context, menu *domain.Menu) error

	Update(ctx context.Context, menu *domain.Menu) error
	Delete(ctx context.Context, id int) error
}
