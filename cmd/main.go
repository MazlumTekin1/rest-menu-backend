package main

//PATH: cmd/main.go

import (
	"log"
	"os"
	"rest-menu-service/internal/adapters/input/http"
	middleware "rest-menu-service/internal/adapters/input/http/middleware"
	gorm "rest-menu-service/internal/adapters/output/gorm"
	"rest-menu-service/internal/application/services"
	"rest-menu-service/internal/infrastructure/config"
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
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		if os.Getenv("GO_ENV") != "production" {
			cfg, err = config.LoadConfig()
			if err != nil {
				log.Fatalf("Failed to load configuration: %v", err)
			}
		} else {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	}

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

	app.Use(func(c *fiber.Ctx) error {
		log.Printf("HTTP Request Received: %s %s", c.Method(), c.Path())
		return c.Next()
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	menuHandler := http.NewMenuHandler(menuService)
	productHandler := http.NewProductHandler(productService)
	categoryHandler := http.NewCategoryHandler(categoryService)

	api := app.Group("/api/v1", middleware.RequireAuth(cfg.JWT.Secret))
	http.SetupMenuRoutes(api, menuHandler)
	http.SetupProductRoutes(api, productHandler)
	http.SetupCategoryRoutes(api, categoryHandler)

	log.Fatal(app.Listen(":4000"))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
