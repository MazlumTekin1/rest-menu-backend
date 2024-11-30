package commands

type CreateCategoryCommand struct {
	Name         string
	MenuID       int
	ImageURL     string
	CreateUserID int
}

type UpdateCategoryCommand struct {
	ID           int
	Name         string
	MenuID       int
	ImageURL     string
	UpdateUserID int
}

type DeleteCategoryCommand struct {
	ID           int
	UpdateUserID int
}
