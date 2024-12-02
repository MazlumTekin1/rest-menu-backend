package domain

//PATH: internal/domain/models.go

import (
	"time"
)

type Menu struct {
	ID           int        `gorm:"column:menu_id;primaryKey;autoIncrement"`
	RestaurantID int        `gorm:"column:restaurant_id;not null"`
	Name         string     `gorm:"column:menu_name;size:100;not null"`
	Categories   []Category `gorm:"foreignKey:MenuID"`
	CreatedDate  time.Time  `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate  time.Time  `gorm:"column:updated_date;autoUpdateTime"`
	IsDeleted    bool       `gorm:"column:is_deleted;default:false"`
	CreateUserID int        `gorm:"column:create_user_id;not null"`
	UpdateUserID int        `gorm:"column:update_user_id;not null"`
}

type Category struct {
	ID           int       `gorm:"column:category_id;primaryKey;autoIncrement"`
	MenuID       int       `gorm:"column:menu_id;not null"`
	Name         string    `gorm:"column:category_name;size:100;not null"`
	ImageURL     string    `gorm:"column:category_image_url"`
	Products     []Product `gorm:"foreignKey:CategoryID"`
	CreatedDate  time.Time `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate  time.Time `gorm:"column:updated_date;autoUpdateTime"`
	IsDeleted    bool      `gorm:"column:is_deleted;default:false"`
	CreateUserID int       `gorm:"column:create_user_id;not null"`
	UpdateUserID int       `gorm:"column:update_user_id;not null"`
}

type Product struct {
	ID           int       `gorm:"column:product_id;primaryKey;autoIncrement"`
	MenuID       int       `gorm:"column:menu_id;not null"`
	CategoryID   int       `gorm:"column:category_id"`
	Name         string    `gorm:"column:product_name;size:100;not null"`
	Price        float64   `gorm:"column:product_price;type:decimal(10,2);not null"`
	Description  string    `gorm:"column:product_description;type:text"`
	ImageURL     string    `gorm:"column:product_image_url"`
	CreatedDate  time.Time `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate  time.Time `gorm:"column:updated_date;autoUpdateTime"`
	IsDeleted    bool      `gorm:"column:is_deleted;default:false"`
	CreateUserID int       `gorm:"column:create_user_id;not null"`
	UpdateUserID int       `gorm:"column:update_user_id;not null"`
}

// TableName methods for GORM
func (Menu) TableName() string {
	return "rest_user.restaurant_menus"
}

func (Category) TableName() string {
	return "rest_user.categories"
}

func (Product) TableName() string {
	return "rest_user.products"
}
