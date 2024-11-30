package dto

type MenuResponse struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	RestaurantID int           `json:"restaurant_id"`
	Categories   []CategoryDTO `json:"categories,omitempty"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
}

type CategoryDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url,omitempty"`
}
