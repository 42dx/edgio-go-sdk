package common

type BaseListResultType[T Property | Env | Variable] struct {
	Total int `mapstructure:"total_items"`
	Items []T `mapstructure:"items"`
}

type FilteredListResultType[T Property | Env | Variable] struct {
	BaseListResultType[T]
	FilteredTotal int
}
