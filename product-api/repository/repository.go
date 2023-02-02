package repository

import (
	"richard-here/soluix/product-api/database"

	"richard-here/soluix/product-api/model"
)

type Repo struct {
	Db database.DBInstance
}

func CreateRepository(db database.DBInstance) Repo {
	return Repo{Db: db}
}

func (r *Repo) GetProducts(pagination Pagination) (*Pagination, error) {
	var products []model.Product

	r.Db.Db.Scopes(paginate(products, &pagination, r.Db.Db)).Find(&products)
	pagination.Rows = products
	return &pagination, nil
}
