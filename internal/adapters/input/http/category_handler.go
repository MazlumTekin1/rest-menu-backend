package http

import (
	"rest-menu-service/internal/application/commands"
	"rest-menu-service/internal/application/queries"
	ports "rest-menu-service/internal/ports/input"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryService ports.CategoryService
}

func NewCategoryHandler(categoryService ports.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

// @Summary Create a new category
// @Description Create a new restaurant category
// @Tags categorys
// @Accept json
// @Produce json
// @Param category body commands.CreateCategoryCommand true "Category Information"
// @Success 201 {object} dto.CategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /categorys [post]
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var cmd commands.CreateCategoryCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	category, err := h.categoryService.CreateCategory(c.Context(), cmd)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create category",
			"details": err.Error(),
		})
	}

	return c.Status(201).JSON(category)
}

// @Summary Update a category
// @Description Update an existing restaurant category
// @Tags categorys
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body commands.UpdateCategoryCommand true "Category Information"
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /categorys/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	var cmd commands.UpdateCategoryCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	if err := h.categoryService.UpdateCategory(c.Context(), cmd); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update category",
			"details": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Category updated successfully",
	})
}

// @Summary Get a category by ID
// @Description Get a restaurant category by its ID
// @Tags categorys
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} dto.CategoryResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /categorys/{id} [get]
func (h *CategoryHandler) GetCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid category ID",
			"details": err.Error(),
		})
	}

	query := queries.GetCategoryQuery{ID: id}
	category, err := h.categoryService.GetCategory(c.Context(), query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to get category",
			"details": err.Error(),
		})
	}

	return c.JSON(category)
}

// @Summary Delete a category
// @Description Soft delete a restaurant category
// @Tags categorys
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 204 "No Content"
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /categorys/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid category ID",
			"details": err.Error(),
		})
	}

	cmd := commands.DeleteCategoryCommand{ID: id}
	if err := h.categoryService.DeleteCategory(c.Context(), cmd); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete category",
			"details": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}

// @Summary List categorys
// @Description Get a list of restaurant categorys
// @Tags categorys
// @Accept json
// @Produce json
// @Param restaurantId query int false "Filter by Menu ID"
// @Success 200 {array} dto.CategoryResponse
// @Failure 500 {object} ErrorResponse
// @Router /categorys [get]
func (h *CategoryHandler) ListCategory(c *fiber.Ctx) error {
	menuID, err := c.ParamsInt("restaurant_id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid restaurant ID",
			"details": err.Error(),
		})
	}

	query := queries.ListCategoriesQuery{MenuID: menuID}
	categorys, err := h.categoryService.ListCategorys(c.Context(), query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to list categorys",
			"details": err.Error(),
		})
	}

	return c.JSON(categorys)
}

// Router setup
func SetupCategoryRoutes(app *fiber.App, handler *CategoryHandler) {
	categoryGroup := app.Group("/api/v1/categories")

	categoryGroup.Post("/", handler.CreateCategory)
	categoryGroup.Put("/:id", handler.UpdateCategory)
	categoryGroup.Get("/:id", handler.GetCategory)
	categoryGroup.Delete("/:id", handler.DeleteCategory)
	categoryGroup.Get("/", handler.ListCategory)
}
