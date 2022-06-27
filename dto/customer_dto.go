package dto

type CustomerCount struct {
	IsStatus int `db:"is_status"`
	Total    int
}

type CustomerAddress struct {
	Address string
	Total   int
}
