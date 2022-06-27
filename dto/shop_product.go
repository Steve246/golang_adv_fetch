package dto

type ShopProductDto struct {
	ShopId       string `db:"id"`
	SiupNumber   string `db:"no_siup"`
	ShopName     string `db:"name"`
	ProductId    string
	ProductName  string
	ProductPrice int `db:"price"`
	ProductStock int `db:"stock"`
}
