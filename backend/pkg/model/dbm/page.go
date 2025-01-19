package dbm

type Page struct {
	PageSize int
	PageNo   int
}

type ModelPage[T any] struct {
	Total int64
	Items []T
}
