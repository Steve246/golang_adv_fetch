package repository

import (
	"errors"

	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

type ShopRepository interface {
	Insert(customer *model.Shop) error
	Update(customer *model.Shop) error
	Delete(id string) error
}

type shopRepository struct {
	db *sqlx.DB
}

func (c *shopRepository) Insert(shop *model.Shop) error {
	_, err := c.db.NamedExec(utils.INSERT_SHOP, shop)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *shopRepository) Update(shop *model.Shop) error {
	_, err := c.db.NamedExec(utils.UPDATE_SHOP, shop)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *shopRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	c.db.MustExec(utils.DELETE_SHOP_PS_HD, id) //test hard delete
	return
}

func NewShopRepository(db *sqlx.DB) ShopRepository {
	cstRepo := new(shopRepository)
	cstRepo.db = db
	return cstRepo
}
