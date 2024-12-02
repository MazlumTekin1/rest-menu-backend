package queries

//PATH: internal/application/queries/menu_queries.go

type GetMenuQuery struct {
	ID int
}

type ListMenusQuery struct {
	RestaurantID int
	Page         int
	PageSize     int
}
