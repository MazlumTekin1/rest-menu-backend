package http

//PATH: internal/adapters/input/http/product_handler.go

import (
	"rest-menu-service/internal/application/commands"
	"rest-menu-service/internal/application/queries"
	ports "rest-menu-service/internal/ports/input"
	errorBody "rest-menu-service/internal/ports/output"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService ports.ProductService
}

func NewProductHandler(productService ports.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// @Summary Create a new product
// @Description Create a new restaurant product
// @Tags products
// @Accept json
// @Produce json
// @Param product body commands.CreateProductCommand true "Product Information"
// @Success 201 {object} dto.ProductResponse
// @Failure 400 {object} ports.ErrorResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /products [post]
// @Security Bearer
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var cmd commands.CreateProductCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	product, err := h.productService.CreateProduct(c.Context(), cmd)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create product",
			"details": err.Error(),
		})
	}

	return c.Status(201).JSON(product)
}

// @Summary Update a product
// @Description Update an existing restaurant product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body commands.UpdateProductCommand true "Product Information"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} ports.ErrorResponse
// @Failure 404 {object} ports.ErrorResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /products/{id} [put]
// @Security Bearer
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var cmd commands.UpdateProductCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	if err := h.productService.UpdateProduct(c.Context(), cmd); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update product",
			"details": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

// @Summary Get a product by ID
// @Description Get a restaurant product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Failure 404 {object} ports.ErrorResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /products/{id} [get]
// @Security Bearer
func (h *ProductHandler) GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid product ID",
			"details": err.Error(),
		})
	}

	query := queries.GetProductQuery{ID: id}
	product, err := h.productService.GetProduct(c.Context(), query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to get product",
			"details": err.Error(),
		})
	}

	return c.JSON(product)
}

// @Summary Delete a product
// @Description Soft delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Failure 404 {object} ports.ErrorResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /products/{id} [delete]
// @Security Bearer
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(errorBody.ErrorResponse{
			Code:    400,
			Message: "Invalid product ID. Error: " + err.Error(),
		})
	}

	cmd := commands.DeleteProductCommand{
		ID:           id,
		UpdateUserID: 1, // Bu değeri actual user'dan almalısınız
	}

	if err := h.productService.DeleteProduct(c.Context(), cmd); err != nil {
		return c.Status(500).JSON(errorBody.ErrorResponse{
			Code:    500,
			Message: "Failed to delete product. Error: " + err.Error(),
		})
	}

	return c.SendStatus(204)
}

// @Summary List products
// @Description Get a list of products with optional filters
// @Tags products
// @Accept json
// @Produce json
// @Param menuId query int false "Filter by Menu ID"
// @Param categoryId query int false "Filter by Category ID"
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Page size (default: 10)"
// @Success 200 {array} dto.ProductResponse
// @Failure 500 {object} ports.ErrorResponse
// @Router /products [get]
// @Security Bearer
func (h *ProductHandler) ListProducts(c *fiber.Ctx) error {
	menuID, _ := strconv.Atoi(c.Query("menuId"))
	categoryID, _ := strconv.Atoi(c.Query("categoryId"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	filter := queries.ListProductsQuery{
		MenuID:     menuID,
		CategoryID: categoryID,
		Page:       page,
		PageSize:   pageSize,
	}

	products, err := h.productService.ListProducts(c.Context(), filter)
	if err != nil {
		return c.Status(500).JSON(errorBody.ErrorResponse{
			Code:    500,
			Message: "Failed to list products" + err.Error(),
		})
	}

	return c.JSON(products)
}

// Router setup
func SetupProductRoutes(api fiber.Router, handler *ProductHandler) {

	api.Post("/products/", handler.CreateProduct)
	api.Put("/products/:id", handler.UpdateProduct)
	api.Get("/products/:id", handler.GetProductById)
	api.Delete("/products/:id", handler.DeleteProduct)
	api.Get("/products/", handler.ListProducts)
}
