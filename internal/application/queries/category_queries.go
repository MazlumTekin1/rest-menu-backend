package queries

//PATH: internal/application/queries/category_queries.go

type GetCategoryQuery struct {
	ID int
}

type ListCategoriesQuery struct {
	MenuID   int
	Page     int
	PageSize int
}
