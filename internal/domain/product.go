package domain

//PATH: internal/domain/product.go

import (
	"errors"
	"rest-menu-service/internal/application/commands"
	"time"
)

// type Product struct {
// 	ID           int
// 	Name         string
// 	RestaurantID int
// 	Categories   []Product
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// 	IsDeleted    bool
// }

func NewProduct(cmd commands.CreateProductCommand) (*Product, error) {
	if cmd.Name == "" {
		return nil, errors.New("product name cannot be empty")
	}

	return &Product{
		Name:         cmd.Name,
		MenuID:       cmd.MenuID,
		CategoryID:   cmd.CategoryID,
		Price:        cmd.Price,
		Description:  cmd.ProductDescription,
		ImageURL:     cmd.ProductImageURL,
		CreatedDate:  time.Now(),
		UpdatedDate:  time.Now(),
		IsDeleted:    false,
		CreateUserID: cmd.CreateUserID,
	}, nil
}
