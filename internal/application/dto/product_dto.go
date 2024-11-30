package dto

type ProductResponse struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Description  string  `json:"description,omitempty"`
	ImageURL     string  `json:"image_url,omitempty"`
	MenuID       int     `json:"menu_id"`
	MenuName     string  `json:"menu_name,omitempty"`
	CategoryID   int     `json:"category_id"`
	CategoryName string  `json:"category_name,omitempty"`
	CreatedDate  string  `json:"created_date"`
	UpdatedDate  string  `json:"updated_date"`
	CreateUserID int     `json:"create_user_id"`
	UpdateUserID int     `json:"update_user_id"`
}
