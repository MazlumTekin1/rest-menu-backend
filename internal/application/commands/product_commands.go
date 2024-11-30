package commands

type CreateProductCommand struct {
	Name               string
	MenuID             int
	CategoryID         int
	Price              float64
	ProductDescription string
	ProductImageURL    string
	CreateUserID       int
}

type UpdateProductCommand struct {
	ID           int
	Name         string
	MenuID       int
	CategoryID   int
	UpdateUserID int
}

type DeleteProductCommand struct {
	ID           int
	UpdateUserID int
}
