package dto

//PATH: internal/application/dto/category_dto.go

type CategoryResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MenuID    int    `json:"menu_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
