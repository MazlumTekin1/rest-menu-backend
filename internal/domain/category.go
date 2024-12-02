package domain

//PATH: internal/domain/category.go

import (
	"errors"
	"time"
)

// type Category struct {
// 	ID           int
// 	Name         string
// 	ImageURL     string
// 	RestaurantID int
// 	Categories   []Category
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// 	IsDeleted    bool
// }

func NewCategory(name string, menuID, createUserID int) (*Category, error) {
	if name == "" {
		return nil, errors.New("category name cannot be empty")
	}

	return &Category{
		Name:         name,
		MenuID:       menuID,
		CreatedDate:  time.Now(),
		UpdatedDate:  time.Now(),
		IsDeleted:    false,
		CreateUserID: createUserID,
	}, nil
}

// func (m *Category) AddCategory(category Category) error {
// 	if category.Name == "" {
// 		return errors.New("category name cannot be empty")
// 	}

// 	m.Categories = append(m.Categories, category)
// 	m.UpdatedAt = time.Now()
// 	return nil
// }
