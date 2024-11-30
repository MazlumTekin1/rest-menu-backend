package ports

import (
	"context"
	"rest-menu-service/internal/application/commands"
	"rest-menu-service/internal/application/dto"
	"rest-menu-service/internal/application/queries"
)

type MenuService interface {
	CreateMenu(ctx context.Context, cmd commands.CreateMenuCommand) (*dto.MenuResponse, error)
	UpdateMenu(ctx context.Context, cmd commands.UpdateMenuCommand) error
	DeleteMenu(ctx context.Context, cmd commands.DeleteMenuCommand) error
	GetMenu(ctx context.Context, query queries.GetMenuQuery) (*dto.MenuResponse, error)
	ListMenus(ctx context.Context, query queries.ListMenusQuery) ([]dto.MenuResponse, error)
}
