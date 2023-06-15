package storage

type RQLRule struct {
	UUID   string      `json:"uuid"`
	Name   string      `json:"name"`
	Script interface{} `json:"script"`
}
