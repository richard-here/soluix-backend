package repository

import (
	"richard-here/soluix/product-api/database"

	localmodel "richard-here/soluix/product-api/model/local"
)

type Repo struct {
	Db database.DBInstance
}

func CreateRepository(db database.DBInstance) Repo {
	return Repo{Db: db}
}

func (r *Repo) GetProducts(pagination Pagination) (*Pagination, error) {
	var products []localmodel.Product

	r.Db.Db.Scopes(paginate(products, &pagination, r.Db.Db)).Find(&products)
	pagination.Rows = products
	return &pagination, nil
}

func (r *Repo) AddProduct(p *localmodel.Product) error {
	err := r.Db.Db.Create(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateProduct(p *localmodel.Product, status bool, retailPrice int) error {
	err := r.Db.Db.First(&p).Error
	if err != nil {
		return err
	}

	p.Status = status
	p.RetailPrice = retailPrice
	err = r.Db.Db.Save(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetProduct(id string) (*localmodel.Product, error) {
	p := new(localmodel.Product)
	err := r.Db.Db.First(&p, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}
