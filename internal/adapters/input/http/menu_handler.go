package http

//PATH: internal/adapters/input/http/menu_handler.go

import (
	"rest-menu-service/internal/application/commands"
	"rest-menu-service/internal/application/queries"
	ports "rest-menu-service/internal/ports/input"
	errorBody "rest-menu-service/internal/ports/output"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MenuHandler struct {
	menuService ports.MenuService
}

func NewMenuHandler(menuService ports.MenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

// @Summary Create a new menu
// @Description Create a new restaurant menu
// @Tags menus
// @Accept json
// @Produce json
// @Param menu body commands.CreateMenuCommand true "Menu Information"
// @Success 201 {object} dto.MenuResponse
// @Failure 400 {object} ports.ErrorResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /menus [post]
// @Security Bearer
func (h *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var cmd commands.CreateMenuCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	menu, err := h.menuService.CreateMenu(c.Context(), cmd)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create menu",
			"details": err.Error(),
		})
	}

	return c.Status(201).JSON(menu)
}

// @Summary Update a menu
// @Description Update an existing restaurant menu
// @Tags menus
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Param menu body commands.UpdateMenuCommand true "Menu Information"
// @Success 200 {object} dto.MenuResponse
// @Failure 400 {object} ports.ErrorResponse
// @Failure 404 {object} ports.ErrorResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /menus/{id} [put]
// @Security Bearer
func (h *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	var cmd commands.UpdateMenuCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	if err := h.menuService.UpdateMenu(c.Context(), cmd); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update menu",
			"details": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Menu updated successfully",
	})
}

// @Summary Get menu by ID
// @Description Get a specific menu by its ID
// @Tags menus
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Success 200 {object} dto.MenuResponse
// @Failure 404 {object} ports.ErrorResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /menus/{id} [get]
// @Security Bearer
func (h *MenuHandler) GetMenuByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(errorBody.ErrorResponse{
			Code:    400,
			Message: "Invalid menu ID. Error: " + err.Error(),
		})
	}

	query := queries.GetMenuQuery{ID: id}
	menu, err := h.menuService.GetMenu(c.Context(), query)
	if err != nil {
		return c.Status(404).JSON(errorBody.ErrorResponse{
			Code:    404,
			Message: "Menu not found. Error: " + err.Error(),
		})
	}

	return c.JSON(menu)
}

// @Summary Delete a menu
// @Description Soft delete a restaurant menu
// @Tags menus
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Success 204 "No Content"
// @Failure 404 {object} ports.ErrorResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /menus/{id} [delete]
// @Security Bearer
func (h *MenuHandler) DeleteMenu(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid menu ID",
			"details": err.Error(),
		})
	}

	cmd := commands.DeleteMenuCommand{ID: id}
	if err := h.menuService.DeleteMenu(c.Context(), cmd); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete menu",
			"details": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Menu deleted successfully",
	})
}

// @Summary List all menus
// @Description Get a list of all menus with optional restaurant filter
// @Tags menus
// @Accept json
// @Produce json
// @Param restaurantId query int false "Filter by Restaurant ID"
// @Success 200 {array} dto.MenuResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /menus [get]
// @Security Bearer
func (h *MenuHandler) ListMenus(c *fiber.Ctx) error {
	restaurantID, _ := strconv.Atoi(c.Query("restaurantId"))

	query := queries.ListMenusQuery{
		RestaurantID: restaurantID,
	}

	menus, err := h.menuService.ListMenus(c.Context(), query)
	if err != nil {
		return c.Status(500).JSON(errorBody.ErrorResponse{
			Code:    500,
			Message: "Failed to list menus. Error: " + err.Error(),
		})
	}

	return c.JSON(menus)
}

func SetupMenuRoutes(api fiber.Router, handler *MenuHandler) {

	api.Post("/menus/", handler.CreateMenu)
	api.Put("/menus/:id", handler.UpdateMenu)
	api.Get("/menus/:id", handler.GetMenuByID)
	api.Delete("/menus/:id", handler.DeleteMenu)
	api.Get("/menus/", handler.ListMenus)
}
