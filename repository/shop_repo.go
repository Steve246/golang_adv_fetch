package repository

import (
	"errors"

	"enigmacamp.com/go-db-fundamnetal/dto"
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

type ShopRepository interface {
	Insert(customer *model.Shop) error
	Update(customer *model.Shop) error
	Delete(id string) error
	GetShopProduct(page int, totalRow int) ([]dto.ShopProductDto, error)
}

type shopRepository struct {
	db *sqlx.DB
}

func (s *shopRepository) GetShopProduct(page int, totalRow int) ([]dto.ShopProductDto, error) {
	limit := totalRow
	offset := limit * (page - 1)

	var shopProducts []dto.ShopProductDto

	rows, err := s.db.Queryx(utils.SELECT_SHOP_WITH_PRODUCT, limit, offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var sp dto.ShopProductDto
		err := rows.StructScan(&sp)
		if err != nil {
			return nil, err
		}
		shopProducts = append(shopProducts, sp)
	}
	return shopProducts, nil

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
