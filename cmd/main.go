package main

import (
	"log"
	"rest-menu-service/internal/adapters/input/http"
	gorm "rest-menu-service/internal/adapters/output/gorm"
	"rest-menu-service/internal/application/services"
	db_gorm "rest-menu-service/internal/infrastructure/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Restaurant Menu API
// @version 1.0
// @description This is a restaurant menu service API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /api/v1
func main() {
	// Setup database with GORM
	db, err := db_gorm.SetupGormDatabase()
	if err != nil {
		log.Fatalf("Could not setup database: %v", err)
	}

	// Setup repositories
	menuRepo := gorm.NewMenuRepository(db)
	productRepo := gorm.NewProductRepository(db)
	categoryRepo := gorm.NewCategoryRepository(db)

	// Setup services
	menuService := services.NewMenuService(menuRepo)
	productService := services.NewProductService(productRepo)
	categoryService := services.NewCategoryService(categoryRepo)

	// Setup Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	// Swagger routes
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Setup API routes
	menuHandler := http.NewMenuHandler(menuService)
	productHandler := http.NewProductHandler(productService)
	categoryHandler := http.NewCategoryHandler(categoryService)

	http.SetupMenuRoutes(app, menuHandler)
	http.SetupProductRoutes(app, productHandler)
	http.SetupCategoryRoutes(app, categoryHandler)

	log.Fatal(app.Listen(":4000"))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	// Retrieve the custom statuscode if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send custom error response
	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
