package repository

import (
	"enigmacamp.com/go-db-fundamnetal/model"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Insert(tr *model.Transaction) error
}

type transactionRepository struct {
	db *sqlx.DB
}

func (t *transactionRepository) Insert(tr *model.Transaction) error {

}
