package queries

//PATH: internal/application/queries/product_queries.go

type ListProductsQuery struct {
	MenuID     int
	CategoryID int
	Page       int
	PageSize   int
}

type GetProductQuery struct {
	ID int
}
