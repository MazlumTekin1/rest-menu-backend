package queries

type GetCategoryQuery struct {
	ID int
}

type ListCategoriesQuery struct {
	MenuID   int
	Page     int
	PageSize int
}
