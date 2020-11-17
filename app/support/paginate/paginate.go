package paginate

type Paginator struct {
	CurrentPage int
	PageSize    int
	Total       int64
}
