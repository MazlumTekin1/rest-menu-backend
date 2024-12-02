package commands

//PATH: internal/application/commands/menu_commands.go
type CreateMenuCommand struct {
	Name         string
	RestaurantID int
	CreateUserID int
}

type UpdateMenuCommand struct {
	ID           int
	Name         string
	RestaurantID int
	UpdateUserID int
}

type DeleteMenuCommand struct {
	ID           int
	UpdateUserID int
}
