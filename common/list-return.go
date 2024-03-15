package common

type BaseListResultType struct {
	Total int   `mapstructure:"total_items"`
	Items []any `mapstructure:"items"`
}
