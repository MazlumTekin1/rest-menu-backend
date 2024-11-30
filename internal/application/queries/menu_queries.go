package queries

type GetMenuQuery struct {
	ID int
}

type ListMenusQuery struct {
	RestaurantID int
	Page         int
	PageSize     int
}
