package types

type APIList struct {
	Total int    `json:"total,omitempty"`
	Data  []*API `json:"data,omitempty"`
}

type API struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type APIListOptions struct {
	Id          string
	Name        string
}
