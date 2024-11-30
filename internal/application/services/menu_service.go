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

type menuService struct {
	menuRepo output.MenuRepository
}

func NewMenuService(menuRepo output.MenuRepository) *menuService {
	return &menuService{
		menuRepo: menuRepo,
	}
}

func (s *menuService) CreateMenu(ctx context.Context, cmd commands.CreateMenuCommand) (*dto.MenuResponse, error) {
	menu, err := domain.NewMenu(cmd.Name, cmd.RestaurantID, cmd.CreateUserID)
	if err != nil {
		return nil, err
	}

	if err := s.menuRepo.Create(ctx, menu); err != nil {
		return nil, err
	}

	return mapToMenuResponse(menu), nil
}

func (s *menuService) UpdateMenu(ctx context.Context, cmd commands.UpdateMenuCommand) error {
	menu, err := s.menuRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	menu.ID = cmd.ID
	menu.RestaurantID = cmd.RestaurantID
	menu.Name = cmd.Name
	menu.UpdatedDate = time.Now()
	menu.UpdateUserID = cmd.UpdateUserID

	return s.menuRepo.Update(ctx, menu)
}

func (s *menuService) DeleteMenu(ctx context.Context, cmd commands.DeleteMenuCommand) error {
	menu, err := s.menuRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	menu.ID = cmd.ID
	menu.UpdatedDate = time.Now()
	menu.UpdateUserID = cmd.UpdateUserID

	return s.menuRepo.Delete(ctx, cmd.ID)
}

func (s *menuService) GetMenu(ctx context.Context, query queries.GetMenuQuery) (*dto.MenuResponse, error) {
	menu, err := s.menuRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	return mapToMenuResponse(menu), nil
}

func (s *menuService) ListMenus(ctx context.Context, query queries.ListMenusQuery) ([]dto.MenuResponse, error) {
	filter := output.MenuFilter{
		RestaurantID: query.RestaurantID,
	}

	menus, err := s.menuRepo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	response := make([]dto.MenuResponse, len(menus))
	for i, menu := range menus {
		response[i] = *mapToMenuResponse(&menu)
	}

	return response, nil
}

func mapToMenuResponse(menu *domain.Menu) *dto.MenuResponse {
	return &dto.MenuResponse{
		ID:           menu.ID,
		Name:         menu.Name,
		RestaurantID: menu.RestaurantID,
		Categories:   mapToCategories(menu.Categories),
		CreatedAt:    menu.CreatedDate.Format(time.RFC3339),
		UpdatedAt:    menu.UpdatedDate.Format(time.RFC3339),
	}
}

func mapToCategories(categories []domain.Category) []dto.CategoryDTO {
	result := make([]dto.CategoryDTO, len(categories))
	for i, cat := range categories {
		result[i] = dto.CategoryDTO{
			ID:       cat.ID,
			Name:     cat.Name,
			ImageURL: cat.ImageURL,
		}
	}
	return result
}
