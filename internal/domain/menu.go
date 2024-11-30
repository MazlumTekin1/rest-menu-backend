package domain

import (
	"errors"
	"time"
)

// type Menu struct {
// 	ID           int
// 	Name         string
// 	RestaurantID int
// 	Categories   []Category
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// 	IsDeleted    bool
// 	CreateUserID int
// }

func NewMenu(name string, restaurantID, createUserID int) (*Menu, error) {
	if name == "" {
		return nil, errors.New("menu name cannot be empty")
	}

	return &Menu{
		Name:         name,
		RestaurantID: restaurantID,
		CreatedDate:  time.Now(),
		UpdatedDate:  time.Now(),
		IsDeleted:    false,
		CreateUserID: createUserID,
	}, nil
}

func (m *Menu) AddCategory(category Category) error {
	if category.Name == "" {
		return errors.New("category name cannot be empty")
	}

	m.Categories = append(m.Categories, category)
	m.UpdatedDate = time.Now()
	return nil
}
