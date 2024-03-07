package common

type BaseListResultType struct {
	Total int           `json:"total_items"`
	Items []interface{} `json:"items"`
}
