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

type productService struct {
	productRepo output.ProductRepository
}

func NewProductService(productRepo output.ProductRepository) *productService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, cmd commands.CreateProductCommand) (*dto.ProductResponse, error) {
	product, err := domain.NewProduct(cmd)
	if err != nil {
		return nil, err
	}

	if err := s.productRepo.Create(ctx, product); err != nil {
		return nil, err
	}

	return mapToProductResponse(product), nil
}

func (s *productService) UpdateProduct(ctx context.Context, cmd commands.UpdateProductCommand) error {
	product, err := s.productRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	product.MenuID = cmd.MenuID
	product.CategoryID = cmd.CategoryID
	product.Name = cmd.Name
	product.UpdatedDate = time.Now()
	product.UpdateUserID = cmd.UpdateUserID

	return s.productRepo.Update(ctx, product)
}

func (s *productService) DeleteProduct(ctx context.Context, cmd commands.DeleteProductCommand) error {
	menu, err := s.productRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	menu.UpdatedDate = time.Now()
	menu.IsDeleted = true
	menu.UpdateUserID = cmd.UpdateUserID

	return s.productRepo.Delete(ctx, cmd.ID, cmd.UpdateUserID)
}

func (s *productService) GetProduct(ctx context.Context, query queries.GetProductQuery) (*dto.ProductResponse, error) {
	product, err := s.productRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	return mapToProductResponse(product), nil
}

func (s *productService) ListProducts(ctx context.Context, query queries.ListProductsQuery) ([]dto.ProductResponse, error) {
	filter := output.ProductFilter{
		MenuID:     query.MenuID,
		CategoryID: query.CategoryID,
		Pagination: output.Pagination{
			Page:     query.Page,
			PageSize: query.PageSize,
		},
	}

	products, err := s.productRepo.ListWithDetails(ctx, filter)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func mapToProductResponse(product *domain.Product) *dto.ProductResponse {
	return &dto.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		MenuID:       product.MenuID,
		CategoryID:   product.CategoryID,
		Price:        product.Price,
		Description:  product.Description,
		ImageURL:     product.ImageURL,
		CreateUserID: product.CreateUserID,
		UpdateUserID: product.UpdateUserID,
		CreatedDate:  product.CreatedDate.Format(time.RFC3339),
		UpdatedDate:  product.UpdatedDate.Format(time.RFC3339),
	}
}
