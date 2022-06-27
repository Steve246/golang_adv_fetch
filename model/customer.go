package model

import "database/sql"

type Customer struct {
	Id       string
	Name     string
	Address  sql.NullString
	Phone    sql.NullString
	Email    string
	Balance  int `db:"saldo"`
	IsStatus int `db:"is_status"`
}

// type CustomerCount struct {
// 	IsStatus int `db:"is_status"`
// 	Total    int
// }

// type CustomerAddress struct {
// 	Address string
// 	Total   int
// }
