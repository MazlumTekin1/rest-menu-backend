package queries

type ListProductsQuery struct {
	MenuID     int
	CategoryID int
	Page       int
	PageSize   int
}

type GetProductQuery struct {
	ID int
}
